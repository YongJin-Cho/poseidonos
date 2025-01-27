#include <gmock/gmock.h>
#include <string>
#include <list>
#include <vector>
#include "src/gc/gc_map_update_request.h"

namespace pos
{
class MockGcMapUpdateRequest : public GcMapUpdateRequest
{
public:
    using GcMapUpdateRequest::GcMapUpdateRequest;
    MOCK_METHOD(bool, Execute, (), (override));
};

} // namespace pos
