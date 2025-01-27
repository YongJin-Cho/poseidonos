#include "src/array/ft/raid1.h"

#include <gtest/gtest.h>

#include <cmath>

#include "src/array_models/dto/partition_physical_size.h"
#include "src/include/address_type.h"

namespace pos
{
TEST(Raid1, Raid1_testIfObjectIsInstantiated)
{
    // Given a set of constructor params
    const PartitionPhysicalSize physicalSize{
        .startLba = 0,
        .blksPerChunk = 10,
        .chunksPerStripe = 5,
        .stripesPerSegment = 20,
        .totalSegments = 100};

    // When
    Raid1 raid1(&physicalSize);

    // Then
    ASSERT_EQ(RaidTypeEnum::RAID1, raid1.GetRaidType());
}

TEST(Raid1, Translate_ifDestinationIsFilledWithStripeIdAndOffset)
{
    // Given
    const PartitionPhysicalSize physicalSize{
        .startLba = 0,
        .blksPerChunk = 10,
        .chunksPerStripe = 5,
        .stripesPerSegment = 20,
        .totalSegments = 100};
    Raid1 raid1(&physicalSize);

    StripeId STRIPE_ID = 1234;
    uint32_t OFFSET = 4567;

    const LogicalBlkAddr src{
        .stripeId = STRIPE_ID,
        .offset = OFFSET};
    FtBlkAddr dest;

    // When
    int actual = raid1.Translate(dest, src);

    // Then
    ASSERT_EQ(0, actual);
    ASSERT_EQ(STRIPE_ID, dest.stripeId);
    ASSERT_EQ(OFFSET, dest.offset);
}

TEST(Raid1, Convert_testIfDestinationIsFilledWithTwoItems)
{
    // Given
    const PartitionPhysicalSize physicalSize{
        .startLba = 0,
        .blksPerChunk = 10,
        .chunksPerStripe = 4,
        .stripesPerSegment = 20,
        .totalSegments = 100};
    Raid1 raid1(&physicalSize);

    StripeId STRIPE_ID = 1234;
    uint32_t OFFSET = 4567;
    uint32_t BACKUP_BLK_CNT = physicalSize.chunksPerStripe / 2 * physicalSize.blksPerChunk; // the semantics defined in Raid1::Raid1()

    const LogicalBlkAddr src{
        .stripeId = STRIPE_ID,
        .offset = OFFSET};
    std::list<BufferEntry> bufferEntries;
    const LogicalWriteEntry srcLogicalWriteEntry = {
        .addr = src,
        .blkCnt = 3,
        .buffers = &bufferEntries};
    list<FtWriteEntry> dest;

    // When
    int actual = raid1.Convert(dest, srcLogicalWriteEntry);

    // Then
    ASSERT_EQ(2, dest.size());

    auto itor = dest.begin();
    FtWriteEntry front = *itor++;
    ASSERT_EQ(STRIPE_ID, front.addr.stripeId);
    ASSERT_EQ(OFFSET, front.addr.offset);

    FtWriteEntry mirror = *itor;
    ASSERT_EQ(STRIPE_ID, mirror.addr.stripeId);
    ASSERT_EQ(OFFSET + BACKUP_BLK_CNT, mirror.addr.offset);
}

TEST(Raid1, GetRebuildGroup_testIfRebuildGroupIsReturnedWhenChunkIndexIsLargerThanMirrorDevCount)
{
    // Given
    const PartitionPhysicalSize physicalSize{
        .startLba = 0,
        .blksPerChunk = 10,
        .chunksPerStripe = 4,
        .stripesPerSegment = 20,
        .totalSegments = 100};
    uint32_t MIRROR_DEV_CNT = physicalSize.chunksPerStripe / 2;
    FtBlkAddr fba;
    int FBA_OFFSET = 1234;
    fba.offset = FBA_OFFSET;
    Raid1 raid1(&physicalSize);

    // When
    list<FtBlkAddr> actual = raid1.GetRebuildGroup(fba);

    // Then
    ASSERT_EQ(1, actual.size());
    int expected_idx = FBA_OFFSET / physicalSize.blksPerChunk;
    int expected_mirror_idx = std::abs((int)(expected_idx - MIRROR_DEV_CNT));
    int expected_offset = expected_mirror_idx * physicalSize.blksPerChunk + (FBA_OFFSET % physicalSize.blksPerChunk);
    FtBlkAddr front = actual.front();
    ASSERT_EQ(expected_offset, front.offset);
}

} // namespace pos
