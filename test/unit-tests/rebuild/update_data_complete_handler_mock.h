#include <gmock/gmock.h>

#include <list>
#include <string>
#include <vector>

#include "src/rebuild/update_data_complete_handler.h"

namespace pos
{
class MockUpdateDataCompleteHandler : public UpdateDataCompleteHandler
{
public:
    using UpdateDataCompleteHandler::UpdateDataCompleteHandler;
    MOCK_METHOD(bool, _DoSpecificJob, (), (override));
};

} // namespace pos
