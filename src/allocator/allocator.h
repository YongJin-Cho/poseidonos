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

#pragma once

#include <string>
#include <vector>

#include "src/allocator/address/allocator_address_info.h"
#include "src/allocator/block_manager/block_manager.h"
#include "src/allocator/context_manager/context_manager.h"
#include "src/allocator/i_allocator_wbt.h"
#include "src/allocator/i_context_replayer.h"
#include "src/allocator/wb_stripe_manager/wbstripe_manager.h"
#include "src/array_models/interface/i_mount_sequence.h"
#include "src/include/meta_const.h"
#include "src/lib/singleton.h"
#include "src/sys_event/volume_event.h"

namespace pos
{
class IArrayInfo;
class IStateControl;

const uint32_t BUFFER_ALLOCATION_SIZE = 2 * 1024 * 1024;
const uint32_t CHUNK_PER_BUFFER_ALLOCATION = BUFFER_ALLOCATION_SIZE / CHUNK_SIZE;
const int BLOCKS_PER_GROUP = 8;

class Allocator : public IAllocatorWbt, public VolumeEvent, public IMountSequence
{
public:
    Allocator(AllocatorAddressInfo* addrInfo, ContextManager* contextManager, BlockManager* blockManager,
              WBStripeManager* wbStripeManager, IArrayInfo* info, IStateControl* iState);
    Allocator(IArrayInfo* info, IStateControl* iState);
    virtual ~Allocator(void);

    virtual int Init(void) override;
    virtual void Dispose(void) override;
    virtual void Shutdown(void) override;
    virtual void Flush(void) override;

    bool VolumeCreated(std::string volName, int volID, uint64_t volSizeBytem, uint64_t maxiops, uint64_t maxbw, std::string arrayName) override { return true; }
    bool VolumeLoaded(std::string name, int id, uint64_t totalSize, uint64_t maxiops, uint64_t maxbw, std::string arrayName) override { return true; }
    bool VolumeUpdated(std::string volName, int volID, uint64_t maxiops, uint64_t maxbw, std::string arrayName) override { return true; }
    bool VolumeMounted(std::string volName, std::string subnqn, int volID, uint64_t volSizeByte, uint64_t maxiops, uint64_t maxbw, std::string arrayName) override { return true; }
    void VolumeDetached(vector<int> volList, std::string arrayName) override {}
    bool VolumeDeleted(std::string volName, int volID, uint64_t volSizeByte, std::string arrayName) override { return true; }
    bool VolumeUnmounted(std::string volName, int volID, std::string arrayName) override;

    void SetGcThreshold(uint32_t inputThreshold) override;
    void SetUrgentThreshold(uint32_t inputThreshold) override;
    int GetMeta(WBTAllocatorMetaType type, std::string fname, MetaFileIntf* file = nullptr) override;
    int SetMeta(WBTAllocatorMetaType type, std::string fname, MetaFileIntf* file = nullptr) override;
    int GetBitmapLayout(std::string fname) override;
    int GetInstantMetaInfo(std::string fname) override;
    void FlushAllUserdataWBT(void) override;

    virtual IBlockAllocator* GetIBlockAllocator(void);
    virtual IWBStripeAllocator* GetIWBStripeAllocator(void);
    IAllocatorWbt* GetIAllocatorWbt(void);
    virtual IContextManager* GetIContextManager(void);
    virtual IContextReplayer* GetIContextReplayer(void);

private:
    void _CreateSubmodules(void);
    void _DeleteSubmodules(void);
    void _RegisterToAllocatorService(void);
    void _UnregisterFromAllocatorService(void);

    AllocatorAddressInfo* addrInfo;
    ContextManager* contextManager;
    BlockManager* blockManager;
    WBStripeManager* wbStripeManager;
    bool isInitialized;

    IArrayInfo* iArrayInfo;
    IStateControl* iStateControl;
};

} // namespace pos
