#include "src/metafs/mim/mfs_io_range_overlap_chker.h"
#include "src/metafs/mim/metafs_io_request.h"
#include <gtest/gtest.h>

namespace pos
{
TEST(MetaFsIoRangeOverlapChker, MetaFsIoRangeOverlapChker_NormalRead)
{
    const MetaLpnType maxLpn = 100;
    const MetaLpnType testLpn = 80;
    const std::string arrayName = "TESTARRAY";
    MetaFsIoRangeOverlapChker* checker = new MetaFsIoRangeOverlapChker();

    checker->Init(maxLpn);

    BitMap* bitmap = checker->GetOutstandingMioMap();
    EXPECT_NE(bitmap, nullptr);

    MetaFsIoRequest req;
    req.fd = 8;
    req.ioMode = MetaIoMode::Async;
    req.reqType = MetaIoRequestType::Read;
    req.targetMediaType = MetaStorageType::SSD;
    req.byteOffsetInFile = 0;
    req.byteSize = MetaFsIoConfig::DEFAULT_META_PAGE_DATA_CHUNK_SIZE;
    req.arrayName = arrayName;
    req.baseMetaLpn = testLpn;

    bool result = checker->IsRangeOverlapConflicted(&req);
    EXPECT_EQ(result, false);

    checker->PushReqToRangeLockMap(&req);

    result = checker->IsRangeOverlapConflicted(&req);
    EXPECT_EQ(result, false);

    checker->FreeLockContext(testLpn, true);

    result = checker->IsRangeOverlapConflicted(&req);
    EXPECT_EQ(result, false);

    delete checker;
}

TEST(MetaFsIoRangeOverlapChker, MetaFsIoRangeOverlapChker_NormalWrite)
{
    const MetaLpnType maxLpn = 100;
    const MetaLpnType testLpn = 80;
    const std::string arrayName = "TESTARRAY";
    MetaFsIoRangeOverlapChker* checker = new MetaFsIoRangeOverlapChker();

    checker->Init(maxLpn);

    BitMap* bitmap = checker->GetOutstandingMioMap();
    EXPECT_NE(bitmap, nullptr);

    MetaFsIoRequest req;
    req.fd = 8;
    req.ioMode = MetaIoMode::Async;
    req.reqType = MetaIoRequestType::Write;
    req.targetMediaType = MetaStorageType::SSD;
    req.byteOffsetInFile = 0;
    req.byteSize = MetaFsIoConfig::DEFAULT_META_PAGE_DATA_CHUNK_SIZE;
    req.arrayName = arrayName;
    req.baseMetaLpn = testLpn;

    bool result = checker->IsRangeOverlapConflicted(&req);
    EXPECT_EQ(result, false);

    checker->PushReqToRangeLockMap(&req);

    result = checker->IsRangeOverlapConflicted(&req);
    EXPECT_EQ(result, true);

    checker->FreeLockContext(testLpn, false);

    result = checker->IsRangeOverlapConflicted(&req);
    EXPECT_EQ(result, false);

    delete checker;
}

} // namespace pos
