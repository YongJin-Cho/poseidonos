#include <gmock/gmock.h>
#include <string>
#include <list>
#include <vector>
#include "src/journal_manager/log_write/log_write_handler.h"

namespace pos
{
class MockLogWriteHandler : public LogWriteHandler
{
public:
    using LogWriteHandler::LogWriteHandler;
    MOCK_METHOD(void, Init, (BufferOffsetAllocator* allocator, JournalLogBuffer* buffer, JournalConfiguration* config), (override));
    MOCK_METHOD(int, AddLog, (LogWriteContext* context), (override));
    MOCK_METHOD(void, LogFilled, (int logGroupId, MapPageList& dirty), (override));
    MOCK_METHOD(void, LogBufferReseted, (int logGroupId), (override));
};

} // namespace pos
