#include <gmock/gmock.h>

#include <list>
#include <string>
#include <vector>

#include "src/journal_manager/statistics/stripe_log_write_status.h"

namespace pos
{
class MockStripeLogWriteStatus : public StripeLogWriteStatus
{
public:
    using StripeLogWriteStatus::StripeLogWriteStatus;
    MOCK_METHOD(void, BlockLogFound, (BlockWriteDoneLog dat), (override));
    MOCK_METHOD(void, StripeLogFound, (StripeMapUpdatedLog dat), (override));
    MOCK_METHOD(void, Print, (), (override));
};

} // namespace pos
