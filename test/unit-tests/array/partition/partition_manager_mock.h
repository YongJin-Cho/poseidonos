#include <gmock/gmock.h>

#include <list>
#include <string>
#include <vector>

#include "src/array/partition/partition_manager.h"

namespace pos
{
class MockPartitionManager : public PartitionManager
{
public:
    using PartitionManager::PartitionManager;
    MOCK_METHOD(const PartitionLogicalSize*, GetSizeInfo, (PartitionType type), (override));
    MOCK_METHOD(int, CreateAll, (vector<ArrayDevice*> buf, vector<ArrayDevice*> data, ArrayInterface* intf), (override));
    MOCK_METHOD(void, DeleteAll, (ArrayInterface * intf), (override));
    MOCK_METHOD(void, FormatMetaPartition, (vector<ArrayDevice*> data, ArrayInterface* intf), (override));
};

} // namespace pos
