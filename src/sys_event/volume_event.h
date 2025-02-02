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

#ifndef __VOLUME_EVENT__
#define __VOLUME_EVENT__

#include <string>
#include <vector>

using namespace std;
namespace pos
{
class VolumeEvent
{
public:
    VolumeEvent(string _tag, string _arrayName)
    {
        tag = _tag;
        arrayName = _arrayName;
    };
    virtual ~VolumeEvent(void);
    string
    Tag()
    {
        return tag;
    }
    void RegisterToPublisher(std::string arrayName);
    void RegisterNvmfTargetToPublisher(std::string arrayName);
    void RemoveFromPublisher(std::string arrayName);
    virtual bool VolumeCreated(string volName, int volID, uint64_t volSizeBytem, uint64_t maxiops, uint64_t maxbw, string arrayName) = 0;
    virtual bool VolumeUpdated(string volName, int volID, uint64_t maxiops, uint64_t maxbw, string arrayName) = 0;
    virtual bool VolumeDeleted(string volName, int volID, uint64_t volSizeByte, string arrayName) = 0;
    virtual bool VolumeMounted(string volName, string subnqn, int volID, uint64_t volSizeByte, uint64_t maxiops, uint64_t maxbw, string arrayName) = 0;
    virtual bool VolumeUnmounted(string volName, int volID, string arrayName) = 0;
    virtual bool VolumeLoaded(string name, int id, uint64_t totalSize, uint64_t maxiops, uint64_t maxbw, string arrayName) = 0;
    virtual void VolumeDetached(vector<int> volList, string arrayName) = 0;

protected:
    string arrayName = "";

private:
    string tag = "";
};
} // namespace pos

#endif // __VOLUME_EVENT__
