/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Intel Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

#include "src/gc/gc_flush_completion.h"

#include "src/spdk_wrapper/free_buffer_pool.h"
#include "src/logger/logger.h"
#include "src/include/backend_event.h"
#include "src/allocator/allocator.h"
#include "src/allocator/wb_stripe_manager/stripe.h"
#include "src/allocator_service/allocator_service.h"
#include "src/array_mgmt/array_manager.h"
#include "src/mapper/mapper.h"
#include "src/mapper_service/mapper_service.h"
#include "src/io/general_io/rba_state_manager.h"
#include "src/io/general_io/rba_state_service.h"
#include "src/io/backend_io/flush_completion.h"
#include "src/gc/copier_meta.h"
#include "src/gc/gc_stripe_manager.h"
#include "src/gc/gc_map_update_request.h"
#include "src/gc/flow_control/flow_control.h"
#include "src/gc/flow_control/flow_control_service.h"
#include "src/array/service/array_service_layer.h"
#include "src/event_scheduler/event_scheduler.h"
#include "Air.h"

#include <list>
#include <string>

namespace pos
{
GcFlushCompletion::GcFlushCompletion(Stripe* stripe, std::string& arrayName, GcStripeManager* gcStripeManager, GcWriteBuffer* dataBuffer)
: Callback(false),
  stripe(stripe),
  arrayName(arrayName),
  gcStripeManager(gcStripeManager),
  dataBuffer(dataBuffer)
{
}

GcFlushCompletion::~GcFlushCompletion(void)
{
}

bool
GcFlushCompletion::_DoSpecificJob(void)
{
    if (nullptr != dataBuffer)
    {
        gcStripeManager->ReturnBuffer(dataBuffer);
        dataBuffer = nullptr;
    }

    RBAStateManager* rbaStateManager =
        RBAStateServiceSingleton::Instance()->GetRBAStateManager(arrayName);

    IArrayInfo* info = ArrayMgr::Instance()->GetArrayInfo(arrayName);
    const PartitionLogicalSize* udSize =
        info->GetSizeInfo(PartitionType::USER_DATA);
    uint32_t totalBlksPerUserStripe = udSize->blksPerStripe;

    BlkAddr rba;
    uint32_t volId;
    std::list<RbaAndSize> rbaList;
    int cnt = 0;

    for (uint32_t i = 0; i < totalBlksPerUserStripe; i++)
    {
        std::tie(rba, volId) = stripe->GetReverseMapEntry(i);
        if (likely(rba != INVALID_RBA))
        {
            RbaAndSize rbaAndSize = {rba * VolumeIo::UNITS_PER_BLOCK,
                BLOCK_SIZE};
            rbaList.push_back(rbaAndSize);
            cnt++;
        }
    }
    std::tie(rba, volId) = stripe->GetReverseMapEntry(0);

    FlowControl* flowControl = FlowControlServiceSingleton::Instance()->GetFlowControl(arrayName);
    int token = flowControl->GetToken(FlowControlType::GC, cnt);
    if (0 >= token)
    {
        return false;
    }

    bool ownershipAcquired = rbaStateManager->AcquireOwnershipRbaList(volId,
            rbaList);
    if (false == ownershipAcquired)
    {
        if (0 < token)
        {
            flowControl->ReturnToken(FlowControlType::GC, token);
        }
        return false;
    }

    airlog("PERF_COPY", "AIR_WRITE", 0, totalBlksPerUserStripe * BLOCK_SIZE);

    EventSmartPtr event(new GcMapUpdateRequest(stripe, arrayName, gcStripeManager));
    stripe->Flush(event);

    return true;
}

} // namespace pos
