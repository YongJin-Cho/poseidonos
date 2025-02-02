# Event

# Event

The following table contains a set of event codes to help interpret the return codes from CLI/REST output.

|Event Code|Event Name|Severity|Description(s)|
|----------|----------|--------|--------------|
| 0 | SUCCESS | - | - |
| 1000 | POS_STARTED | - | - |
| 1001 | POS_TERMINATED | - | - |
| 1002 | SERVER_READY | INFO | CLI server initialized successfully. |
| 1003 | CLIENT_CONNECTED | INFO | new client {} is connected |
| 1004 | CLIENT_DISCONNECTED | INFO | client {} is disconnected |
| 1005 | MSG_RECEIVED | INFO | msg received:{} |
| 1006 | MSG_SENT | INFO | msg sent:{} |
| 1010 | SERVER_TRY_EXIT | INFO | server try to exit |
| 1011 | SERVER_THREAD_JOINED | INFO | server thread joined |
| 1012 | REUSE_ADDR_ENABLED | INFO | socket opt - reuseaddr enabled |
| 1013 | REUSE_ADDR_FAILED | WARN | failed to enable reuse addr |
| 1014 | SOCK_CREATE_FAILED | ERROR | socket creation failed |
| 1015 | SOCK_BIND_FAILED | ERROR | socket biding failed |
| 1016 | SOCK_LISTEN_FAILED | ERROR | socket listen failed |
| 1017 | EPOLL_CREATE_FAILED | ERROR | failed to create epoll |
| 1018 | SOCK_ACCEPT_FAILED | WARN | failed to accept socket |
| 1019 | MAX_CLIENT_ERROR | WARN | max client exceed |
| 1020 | MSG_SEND_FAILED | ERROR | failed to send msg:{} |
| 1021 | MSG_RECEIVE_EXCEPTION | - | exception on msg receiving:{} |
| 1022 | TIMED_OUT | INFO | TIMED OUT |
| 1023 | INVALID_PARAM | - | volume name or newname is not entered |
| 1030 | POS_BUSY | - | POS IS BUSY |
| 1101 | STATE_CONTEXT_UPDATED | DEBUG | statecontext added - {} |
| 1102 | STATE_CHANGED | INFO | Array state changed to {} |
| 1103 | STATE_CONTROL_ADDED | INFO | statecontrol of array:{} is added |
| 1104 | STATE_CONTROL_REMOVED | INFO | statecontrol of array:{} is removed |
| 1105 | STATE_CONTROL_DEBUG | INFO | statecontrol of array:{} is nullptr |
| 1234 | SYSTEM_RECOVERY | TRACE | progress report: [ |
| 2000 | VOL_CREATED | INFO | Volume name:{} size:{} created |
| 2001 | VOL_DELETED | - | - |
| 2002 | VOL_MOUNTED | INFO | Volume mounted name: {} |
| 2003 | VOL_UNMOUNTED | INFO | Volume unmounted name: {} |
| 2004 | VOL_ADDED | DEBUG | Volume added to the list, VOL_CNT: {}, VOL_ID: {} |
| 2005 | VOL_REMOVED | INFO | Volume removed from the list VOL_CNT {} |
| 2010 | VOL_NOT_EXIST | WARN | The requested volume does not exist |
| 2011 | VOL_CNT_EXCEEDED | WARN | Excced maximum number of volumes |
| 2012 | VOL_SIZE_EXCEEDED | WARN | The requested volume size is larger than the remaining space |
| 2020 | VOL_NAME_TOO_SHORT | WARN | Volume name must be at least {} characters |
| 2021 | VOL_NAME_TOO_LONG | WARN | Volume name must be less or equal than {} characters |
| 2022 | VOL_NAME_DUPLICATED | WARN | Volume name is duplicated |
| 2023 | VOL_NAME_NOT_ALLOWED | WARN | Blank cannot be placed at the beginning or end of a volume name |
| 2030 | VOL_SIZE_TOO_BIG | - | - |
| 2031 | VOL_SIZE_TOO_SMALL | WARN | The requested volume size is less than already used |
| 2032 | VOL_SIZE_NOT_ALIGNED | WARN | The requested size, {} is not aligned to MB |
| 2040 | VOL_ALD_MOUNTED | WARN | The volume already mounted: {} |
| 2041 | VOL_ALD_UNMOUNTED | WARN | The volume already unmounted: {} |
| 2050 | DEL_MOUNTED_VOL | WARN | Unable to delete mounted volume |
| 2060 | META_CREATE_FAIL | ERROR | Fail to create meta file |
| 2061 | META_OPEN_FAIL | ERROR | Fail to open volume meta |
| 2062 | META_READ_FAIL | ERROR | Fail to read volume meta |
| 2063 | META_WRITE_FAIL | ERROR | Fail to write volume meta |
| 2064 | META_CONTENT_BROKEN | ERROR | Volume meta broken {} |
| 2070 | MEM_ALLOC_FAIL | ERROR | Fail to allocate memory |
| 2071 | INVALID_INDEX | ERROR | Invalid index error |
| 2072 | VOLID_ALLOC_FAIL | - | - |
| 2073 | VOL_SAMEID_EXIST | ERROR | The same ID volume exists |
| 2080 | OUT_OF_QOS_RANGE | - | - |
| 2090 | ARRAY_NOT_MOUNTED | WARN | Array is not mounted |
| 2091 | SYSTEM_FAULT | WARN | Cannot be performed in the STOP state |
| 2092 | DONE_WITH_ERROR | WARN | Failed to unmount volume during rollback mount: {} |
| 2100 | SUBSYSTEM_NOT_CREATED | WARN | No subsystem:{} was created to attach the volume |
| 2101 | VOL_ALD_SET_SUBNQN | INFO | The volume already has set subsystem {}, replace to {} |
| 2102 | CANNOT_EXTEND_NSID | WARN | Can't extend more nsid to subsystem:{} |
| 2103 | VOL_UNMOUNT_FAIL | ERROR | Nvmf internal error occurred during volume unmount |
| 2104 | VOL_DETACH_FAIL | ERROR | Detach volume failed due to internal error or unmount timeout. Only some of them might be unmounted |
| 2110 | VOL_DATA_SIZE_TOO_BIG | ERROR | Volume meta write buffer overflows |
| 2111 | VOL_DATA_SIZE_TOO_SMALL | - | - |
| 2112 | INVALID_VOL_ID_ERROR | WARN | invalid vol id. vol name : {} |
| 2300 | MBR_ALLOCATE_MEMORY | INFO | MBR allocate memory done |
| 2301 | MBR_CREATE_AND_WRITE | - | - |
| 2302 | MBR_ADD_AND_WRITE | - | - |
| 2303 | MBR_WRITE_DONE | DEBUG | write mbr data |
| 2304 | MBR_READ_DONE | DEBUG | read mbr data |
| 2305 | MBR_PARITY_CHECK | WARN | mbr parity check fail, mbr is reset |
| 2306 | MBR_SYSTEM_UUID_CHECK | INFO | mbr system uuid check  |
| 2310 | MBR_DEVICE_NOT_FOUND | WARN | device not found |
| 2311 | MBR_PARITY_CHECK_FAILED | WARN | mbr parity check fail, mbr is reset |
| 2312 | MBR_SYSTEM_UUID_CHECK_FAILED | WARN | mbr system uuid check fail |
| 2313 | MBR_DATA_NOT_FOUND | WARN | LISTARRAY |
| 2314 | MBR_KEY_NOT_FOUND | - | - |
| 2315 | MBR_ABR_NOT_FOUND | WARN | No array found with arrayName :{} |
| 2316 | MBR_DELETE_ABR_FAILED | - | - |
| 2317 | MBR_MAX_ARRAY_CNT_EXCEED | - | - |
| 2318 | MBR_ABR_ALREADY_EXIST | - | - |
| 2319 | MBR_WRONG_ARRAY_VALID_FLAG | ERROR | Array Valid Flag doesn't match with Index Map |
| 2320 | MBR_WRONG_ARRAY_INDEX_MAP | ERROR | Array index map doesn't match with previous condition |
| 2321 | MBR_DEVICE_ALREADY_IN_ARRAY | DEBUG | Device {} is already in array {} |
| 2322 | MBR_ABR_LIST_SUCCESS | DEBUG | Found {} arrays from abr list |
| 2330 | MBR_ALLOCATE_MEMORY_FAILED | - | - |
| 2331 | MBR_DATA_SIZE_ERROR | - | - |
| 2332 | MBR_FORMAT_ERROR | - | - |
| 2333 | MBR_GET_SYSTEM_UUID_FAILED | ERROR | mbr get system uuid failed |
| 2390 | MBR_TIME_CALC_ERROR | ERROR | Time calculation error |
| 2399 | MBR_DEBUG_MSG | DEBUG | Inserted {} to array {} |
| 2500 | ARRAY_STATE_ONLINE | ERROR | Failed to load array, already mounted |
| 2501 | ARRAY_STATE_OFFLINE | ERROR | Failed to unmount array, not mounted |
| 2502 | ARRAY_STATE_EXIST | WARN | Array force-mounted with degraded state |
| 2503 | ARRAY_STATE_NOT_EXIST | WARN | Array {} is not exist |
| 2504 | ARRAY_DEVICE_NOT_FOUND | WARN | Device not found |
| 2505 | ARRAY_DEVICE_COUNT_ERROR | - | - |
| 2506 | ARRAY_NVM_CAPACITY_ERROR | WARN | NVM device size error |
| 2507 | ARRAY_SSD_CAPACITY_ERROR | ERROR | SSD capacity is not valid. Valid capacity is from {}GB to {}TB |
| 2508 | ARRAY_SSD_SAME_CAPACITY_ERROR | WARN | SSDs must be the same sizes |
| 2509 | ARRAY_MBR_READ_ERROR | - | - |
| 2510 | ARRAY_MBR_WRITE_ERROR | ERROR | MBR Write Error |
| 2511 | ARRAY_FT_METHOD_ERROR | - | - |
| 2512 | ARRAY_WRONG_FT_METHOD | - | RAID1 |
| 2513 | ARRAY_PARTITION_CREATION_ERROR | ERROR | Failed to create partition \ |
| 2514 | ARRAY_PARTITION_LOAD_ERROR | ERROR | Failed to create partition \ |
| 2515 | ARRAY_WRONG_EVENT_ID | - | - |
| 2516 | ARRAY_INVALID_ADDRESS_ERROR | ERROR | Invalid Address Error |
| 2517 | ARRAY_BROKEN_ERROR | ERROR | Array cannot be configured due to faults of {} data devices |
| 2518 | ARRAY_DEVICE_TYPE_ERROR | - | - |
| 2519 | ARRAY_PARTITION_TRIM | DEBUG | Try to trim from {} for {} on {} |
| 2520 | ARRAY_DEVICE_DETACHED | INFO | Spare-device {} is detached from array {} |
| 2521 | ARRAY_NO_REMAINING_SPARE | WARN | No remaining spare device |
| 2522 | ARRAY_WRONG_NAME | - | LISTVOLUME |
| 2523 | ARRAY_NEED_MOUNT | WARN | Failed to add spare. Spare cannot be added without Array mounted |
| 2524 | ARRAY_CNT_EXCEEDED | DEBUG | ArrayManager cnt exceeded. current: {} |
| 2525 | ARRAY_ALREADY_EXIST | DEBUG | - |
| 2526 | ARRAY_NOT_FOUND | WARN | No array found for device serial number {}. DeviceDetached event will be ignored. |
| 2527 | ARRAY_SHUTDOWN_TAKES_TOO_LONG | - | - |
| 2528 | ARRAY_LOAD_FAIL | - | ARRAYINFO |
| 2529 | ARRAY_SERVICE_REGISTRATION_FAIL | - | - |
| 2530 | ARRAY_DEVICE_ADDED | INFO | Spare device was successfully added to array {} |
| 2531 | ARRAY_DEVICE_REMOVED | INFO | The spare device {} removed from array {} |
| 2532 | ARRAY_DEVICE_ADD_FAIL | WARN | set nvm device without reset previous nvm |
| 2533 | ARRAY_DEVICE_REMOVE_FAIL | - | - |
| 2534 | ARRAY_DEVICE_CLEARED | INFO | the array devices are cleared |
| 2535 | ARRAY_DEVICE_REPLACED | INFO | Faulty device is replaced to the spare {} |
| 2536 | ARRAY_DEVICE_WRONG_NAME | WARN | URAM's name does not conform to the naming rule\n |
| 2537 | ARRAY_DEVICE_ALREADY_ADDED | - | - |
| 2538 | ARRAY_DEVICE_NVM_NOT_FOUND | ERROR | Failed to load array {}, check uram creation or pmem state |
| 2539 | ARRAY_DEVICE_NEED_REBUILD | - | - |
| 2540 | ARRAY_ARRAY_INFO_FOUND | DEBUG | Found array '{}' in state '{}' |
| 2541 | ARRAY_NO_ARRAY_INFO | ERROR | No array info for array '{}' |
| 2542 | ARRAY_DEVICE_REBUILD_STATE | DEBUG | Rebuilding device found {} |
| 2543 | ARRAY_INVALID_INDEX | - | - |
| 2550 | ARRAY_STATE_CHANGED | INFO | Array state changed to {} |
| 2551 | ARRAY_STATE_EXIST_NORMAL | - | - |
| 2552 | ARRAY_STATE_EXIST_DEGRADED | WARN | Array force-mounted with degraded state |
| 2553 | ARRAY_STATE_BROKEN | ERROR | Failed to mount array. Array is broken |
| 2554 | ARRAY_STATE_NORMAL | - | - |
| 2555 | ARRAY_STATE_DEGRADED | - | - |
| 2556 | ARRAY_STATE_REBUILDING | ERROR | Failed to unmount array, array is being rebuilt |
| 2557 | ARRAY_STATE_STOP | - | - |
| 2558 | ARRAY_STATE_CHANGE_ERROR | WARN | Failed to change array state, curret state is {} |
| 2560 | ARRAY_NAME_TOO_SHORT | WARN | Array name must be at least {} characters |
| 2561 | ARRAY_NAME_TOO_LONG | WARN | Array name must be less or equal than {} characters |
| 2562 | ARRAY_NAME_NOT_ALLOWED | WARN | Blank cannot be placed at the beginning or end of a array name |
| 2580 | MOUNTED_ARRAY_EXISTS | ERROR | Failed to exit system. Array '{}' is still mounted with state '{}' |
| 2581 | ARRAY_ALD_UNMOUNTED | ERROR | Failed to unmount system. Curr. state:{} |
| 2582 | ARRAY_MOUNT_PRIORITY_ERROR | DEBUG | Initializing the first mount sequence for {} |
| 2583 | ARRAY_UNMOUNT_PRIORITY_ERROR | DEBUG | Detaching volumes for {} |
| 2584 | ARRAY_UNMOUNTING | INFO | Start flushing all active stripes |
| 2585 | ARRAY_MOUNTING | INFO | Start initializing mapper of array {} |
| 2586 | ARRAY_ALD_MOUNTED | - | - |
| 2590 | ARRAY_DEBUG_MSG | INFO | Array {} loaded successfully |
| 2591 | ARRAY_COMPONENTS_LEAK | WARN | Memory leakage found for ArrayMountSequence for  |
| 2592 | ARRAY_COMPONENTS_DEBUG_MSG | DEBUG | Deleting array component for {} |
| 2593 | ARRAY_MOUNTSEQ_DEBUG_MSG | DEBUG | Entering ArrayMountSequence.Mount for {} |
| 2594 | ARRAY_MANAGER_DEBUG_MSG | DEBUG | Ignoring DeviceManager's subscription to DeviceEvent. It's not okay if it's for prod code (not unit test) |
| 2700 | TRANSLATOR_DEBUG_MSG | INFO | IOTranslator::Register, array:{} size:{} |
| 2701 | TRANSLATOR_NOT_EXIST | ERROR | IOTranslator::Translate ERROR, array:{} part:{} |
| 2710 | RECOVER_DEBUG_MSG | INFO | IORecover::Register, array:{} size:{} |
| 2711 | RECOVER_IS_RETRY | ERROR | RebuildRead::GetRecoverMethod IsRetry |
| 2712 | RECOVER_FAILED | - | - |
| 2713 | RECOVER_INVALID_LBA | - | - |
| 2720 | DEV_CHECKER_DEBUG_MSG | INFO | IODeviceChecker::Register, array:{} |
| 2730 | LOCKER_DEBUG_MSG | ERROR | Locker not found, fatal error if array is not broken |
| 2800 | REBUILD_DEBUG_MSG | DEBUG | normallocker is now changing state. Using stripe {} is refused |
| 2802 | REBUILD_STOPPED | WARN | Partition {} (RAID1) rebuilding stopped |
| 2803 | REBUILD_FAILED | WARN | Partition {} (RAID1) rebuilding failed |
| 2804 | REBUILD_PROGRESS | DEBUG | [0], {} |
| 2805 | REBUILD_PROGRESS_DETAIL | DEBUG | array:{}, id:{}, done:{} |
| 2806 | REBUILD_STRIPE_LOCK | - | - |
| 2807 | REBUILD_STRIPE_UNLOCK | - | - |
| 2810 | REBUILD_RESULT_PASS | INFO | array {} rebuild completed sucessfully |
| 2811 | REBUILD_RESULT_FAILED | - | - |
| 2812 | REBUILD_RESULT_CANCELLED | WARN | array {} rebuild cancelled |
| 2813 | REBUILD_TRIGGER_FAIL | WARN | Failed to trigger rebuild. Current array state is not rebuildable |
| 2900 | CONFIG_FILE_READ_DONE | - | - |
| 2910 | CONFIG_REQUEST_CONFIG_TYPE_ERROR | - | - |
| 2911 | CONFIG_REQUEST_KEY_ERROR | - | - |
| 2912 | CONFIG_REQUEST_MODULE_ERROR | - | - |
| 2920 | CONFIG_FILE_OPEN_FAIL | ERROR | Cannot open uram config file |
| 2921 | CONFIG_FILE_SIZE_ERROR | - | - |
| 2922 | CONFIG_FILE_READ_ERROR | - | - |
| 2923 | CONFIG_FILE_FORMAT_ERROR | - | - |
| 2924 | CONFIG_JSON_DOC_IS_NOT_OBJECT | - | - |
| 2925 | CONFIG_VALUE_TYPE_ERROR | - | - |
| 3000 | JOURNAL_MANAGER_INITIALIZED | INFO | Journal manager is initialized to status {} |
| 3001 | JOURNAL_MANAGER_NOT_INITIALIZED | ERROR | Journal manager accessed without initialization |
| 3002 | JOURNAL_CONFIGURATION | INFO | Failed to read journal enablement from config file |
| 3003 | JOURNAL_NOT_READY | - | - |
| 3004 | JOURNAL_ALREADY_EXIST | INFO | Journal service for array {} is already registered |
| 3010 | JOURNAL_LOG_BUFFER_CREATED | INFO | Log buffer is created |
| 3011 | JOURNAL_LOG_BUFFER_CREATE_FAILED | ERROR | Log buffer already exists |
| 3012 | JOURNAL_LOG_BUFFER_LOADED | INFO | Journal log buffer is loaded |
| 3013 | JOURNAL_LOG_BUFFER_RESET | ERROR | Failed to reset journal log buffer |
| 3014 | JOURNAL_LOG_BUFFER_RESET_FAILED | ERROR | Failed to reset journal log buffer |
| 3015 | JOURNAL_LOG_BUFFER_OPEN_FAILED | ERROR | Failed to open log buffer |
| 3016 | JOURNAL_LOG_BUFFER_CLOSE_FAILED | ERROR | Failed to close journal log buffer |
| 3017 | JOURNAL_LOG_BUFFER_READ_FAILED | ERROR | Failed to read log buffer |
| 3020 | JOURNAL_LOG_WRITE_FAILED | ERROR | Failed to write journal log |
| 3021 | JOURNAL_INVALID_SIZE_LOG_REQUESTED | ERROR | Requested log size is bigger than meta page |
| 3022 | ADD_TO_JOURNAL_WAITING_LIST | - | - |
| 3023 | JOURNAL_CALLBACK_FAILED | - | - |
| 3024 | JOURNAL_NO_LOG_BUFFER_AVAILABLE | WARN | No log buffer available for journal |
| 3030 | JOURNAL_LOG_GROUP_FULL | DEBUG | Log group id {} is added to full log group |
| 3031 | JOURNAL_FLUSH_LOG_GROUP | DEBUG | Flush next log group {} |
| 3032 | JOURNAL_CHECKPOINT_STARTED | INFO | Checkpoint started with {} maps to flush |
| 3033 | JOURNAL_CHECKPOINT_STATUS | INFO | Allocator meta flush completed |
| 3034 | JOURNAL_CHECKPOINT_COMPLETED | INFO | Checkpoint completed |
| 3035 | JOURNAL_CHECKPOINT_FAILED | ERROR | Failed to start flushing dirty map pages |
| 3036 | JOUNRAL_WRITE_LOG_GROUP_FOOTER | DEBUG | Failed to write log group footer |
| 3040 | JOURNAL_HANDLE_VOLUME_DELETION | DEBUG | Start volume delete event handler (volume id {}) |
| 3050 | JOURNAL_REPLAY_STARTED | INFO | Journal replay started |
| 3051 | JOURNAL_REPLAY_STATUS | DEBUG | SegmentId:{} is allocated |
| 3052 | JOURNAL_REPLAY_STOPPED | DEBUG | No logs to replay. Stop replaying |
| 3053 | JOURNAL_REPLAY_FAILED | CRITICAL | Journal replay failed |
| 3054 | JOURNAL_REPLAY_COMPLETED | INFO | Journal replay completed |
| 3055 | JOURNAL_INVALID_LOG_FOUND | ERROR | Unknown type of log is found |
| 3060 | JOURNAL_REPLAY_STRIPE_FLUSH | DEBUG | Failed to reconstruct active stripe, wb lsid  |
| 3061 | JOURNAL_REPLAY_STRIPE_FLUSH_FAILED | DEBUG | Failed to reconstruct active stripe, wb lsid  |
| 3062 | JOURNAL_REPLAY_WB_TAIL | - | - |
| 3063 | JOURNAL_REPLAY_USER_STRIPE_TAIL | DEBUG | [Replay] SSD LSID is updated to  |
| 3064 | JOURNAL_REPLAY_VOLUME_EVENT | DEBUG | [Replay] {} block log of volume {} is skipped |
| 3100 | MAPPER_SUCCESS | INFO | AsyncStore() for volumeId:{} Started |
| 3101 | MAP_FLUSH_STARTED | DEBUG | FlushDirtyPagesGiven mapId:{} started |
| 3102 | MAP_FLUSH_IN_PROGRESS | - | - |
| 3103 | MAP_FLUSH_ONGOING | DEBUG | Issue mpage flush, startMpage {} numMpages {} |
| 3104 | MAP_FLUSH_COMPLETED | INFO | mapId:{} Flushed @MapAsyncFlushed |
| 3105 | MAP_LOAD_STARTED | INFO | Header Async Loading Started, mapId:{} |
| 3106 | MAP_LOAD_ONGOING | INFO | fileName:{} header load success, valid:{} / total:{} |
| 3107 | MAP_LOAD_COMPLETED | INFO | No mpage to Load, so VSAMap Async Load Finished, volID:{} @_VSAMapFileAsyncLoad |
| 3108 | MPAGE_NULLPTR | ERROR | mpage is nullptr when vsid:{} |
| 3109 | MPAGE_ALREADY_EXIST | ERROR | mpage exists but tried to allocate, pageNr:{} Mpage.data:{0:x} |
| 3110 | MPAGE_MEMORY_ALLOC_FAILURE | ERROR | posix_memalign() failure, ret:{} |
| 3111 | VSAMAP_NULL_PTR | - | - |
| 3112 | VSAMAP_SET_FAILURE | ERROR | VSAMap set failure, volumeId:{}  targetRba:{}  targetVsa.sid:{}  targetVsa.offset:{} |
| 3113 | STRIPEMAP_SET_FAILURE | ERROR | StripeMap set failure, vsid:{}  lsid:{}  loc:{} |
| 3114 | MAPCONTENT_HEADER_NOT_INITIALIZED | - | - |
| 3115 | INVALID_VOLUME_ID | ERROR | VolumeId is invalid |
| 3116 | REVMAP_GET_MFS_ALIGNED_IOSIZE_FAILURE | CRITICAL | MFS returned failure value |
| 3117 | REVMAP_PACK_ALREADY_LINKED | ERROR | Stripe object for wbLsid:{} is already linked to ReverseMapPack, GcStripe:{} |
| 3118 | REVMAP_NOT_LINKED_PACK | ERROR | ReverseMapPack for wbLsid:{} is not linked but tried to unlink |
| 3119 | REVMAP_FILE_SIZE | INFO | fileSizePerStripe:{}  maxVsid:{}  fileSize:{} for RevMapWhole |
| 3120 | REVMAP_VOLUME_ID_FOUND | - | - |
| 3121 | REVMAP_VOLUME_ID_NOT_FOUND | - | - |
| 3122 | REVMAP_MFS_IO_ERROR | ERROR | RevMap MFS IO Error:{}  IoDirection:{} |
| 3123 | MAP_UPDATE_HANDLER_EVENT_ALLOCATE_FAIL | ERROR | Failed to allocate flush wrapup event |
| 3124 | TRIED_TO_LOAD_WITHOUT_MFSFILE | - | - |
| 3125 | NO_BLOCKMAP_MFS_FILE | DEBUG | No MFS filename:{} for volName:{} @VolumeDeleted |
| 3126 | VSAMAP_HEADER_LOAD_FAILURE | WARN | Getting volume Size failed, volumeID:{}  volSizebyVolMgr:{} |
| 3127 | VSAMAP_LOAD_FAILURE | ERROR | VolumeId:{} Internal Mount Request Failed |
| 3128 | VSAMAP_NOT_LOADED | - | - |
| 3129 | VSAMAP_UNMOUNT_FAILURE | WARN | volumeID:{} is Not FG_MOUNTED @VolumeUnmounted |
| 3130 | VSAMAP_INVALIDATE_ALLBLKS_FAILURE | WARN | VSAMap Invalidate all blocks Failed, volumeID:{} @VolumeDeleted |
| 3131 | VSAMAP_UNLOAD_FAILURE | - | - |
| 3132 | VSAMAP_STORE_FAILURE | ERROR | VSAMap SyncStore failed, volumeId:{} |
| 3133 | STRIPEMAP_STORE_FAILURE | ERROR | AsyncStore() of stripeMapManager Failed, Check logs |
| 3134 | MFS_SYNCIO_ERROR | ERROR | AppendIO Error, retMFS:{}  fd:{} |
| 3135 | MFS_ASYNCIO_ERROR | ERROR | MFS AsyncIO error, ioError:{}  mpageNum:{} |
| 3136 | VSAMAP_NOT_ACCESSIBLE | WARN | VolumeId:{} is not accessible, maybe unmounted |
| 3137 | DELETE_VOLUME | INFO | Enable Internal VsaMap Access volumeId:{} |
| 3138 | WRONG_MAP_ID | - | - |
| 3139 | MAPPER_ALREADY_EXIST | ERROR | Allocator for array {} is already registered |
| 3140 | MAPPER_FAILED | ERROR | AsyncStore() for volumeId:{} Failed |
| 3150 | ALLOCATOR_META_ARCHIVE_STORE | DEBUG | Ready to flush RebuildCtx file:{}, numTargetSegments:{} |
| 3151 | ALLOCATOR_META_ARCHIVE_LOAD | DEBUG | RebuildCtx file loaded, segmentCount:{} |
| 3152 | ALLOCATOR_META_ARCHIVE_STORE_REBUILD_SEGMENT | DEBUG | Ready to flush RebuildCtx file:{}, numTargetSegments:{} |
| 3153 | ALLOCATOR_META_ARCHIVE_LOAD_REBUILD_SEGMENT | DEBUG | RebuildCtx file loaded, segmentCount:{} |
| 3154 | ALLOCATOR_META_ARCHIVE_FLUSH_IN_PROGRESS | ERROR | - |
| 3155 | ALLOCATOR_MAKE_REBUILD_TARGET | ERROR | Failed to load RebuildCtx, segmentId:{} is already in set |
| 3156 | ALLOCATOR_MAKE_REBUILD_TARGET_FAILURE | ERROR | Failed to load RebuildCtx, segmentId:{} is already in set |
| 3157 | ALLOCATOR_TARGET_SEGMENT_FREED | INFO | Skip Rebuilding segmentId:{}, Already Freed |
| 3158 | ALLOCATOR_REBUILD_TARGET_SET_NOT_EMPTY | WARN | targetSegmentList is NOT empty! |
| 3159 | ALLOCATOR_REBUILD_TARGET_SET_EMPTY | INFO | Rebuild was already done or not happen |
| 3160 | ALLOCATOR_NO_FREE_WB_STRIPE | - | - |
| 3161 | ALLOCATOR_REBUILDING_SEGMENT | DEBUG | segmentId:{} is already rebuild target! |
| 3162 | ALLOCATOR_REBUILT_SEGMENT | - | - |
| 3163 | ALLOCATOR_NO_FREE_SEGMENT | ERROR | Failed to allocate segment, free segment count:{} |
| 3165 | ALLOCATOR_CANNOT_ALLOCATE_USER_BLOCK | - | - |
| 3166 | ALLOCATOR_CANNOT_ALLOCATE_STRIPE | ERROR | failed to allocate gc stripe! |
| 3167 | ALLOCATOR_CANNOT_LINK_REVERSE_MAP | ERROR | failed to link ReverseMap to allocate gc stripe! |
| 3168 | ALLOCATOR_STRIPE_WITHOUT_REVERSEMAP | ERROR | Stripe object for wbLsid:{} is not linked to reversemap but tried to Unlink, GcStripe:{} |
| 3169 | ALLOCATOR_RECONSTRUCT_STRIPE | DEBUG | Stripe (vsid {}, wbLsid {}, blockCount {}, remainingBlks {}) is reconstructed |
| 3170 | ALLOCATOR_REPLAYED_STRIPE_IS_FULL | DEBUG | Stripe (vsid {}, wbLsid {}) is waiting to be flushed |
| 3171 | ALLOCATOR_REPLAY_SEGMENT_STATUS | - | - |
| 3172 | ALLOCATOR_TRIGGER_FLUSH | DEBUG | Request stripe flush, vsid {} lsid {} remaining {} |
| 3173 | REVMAP_RECONSTRUCT_FOUND_RBA | INFO | [RMR]START, volumeId:{}  wbLsid:{}  vsid:{}  blockCount:{} |
| 3174 | REVMAP_RECONSTRUCT_NOT_FOUND_RBA | INFO | There was no vsa map entry for some blocks |
| 3175 | GET_VOLUMESIZE_FAILURE | WARN | [RMR]GetVolumeSize failure, volumeId:{} |
| 3177 | PICKUP_ACTIVE_STRIPE | INFO | Picked Active Stripe: index:{}  wbLsid:{}  vsid:{}  remaining:{} |
| 3178 | SEGMENT_WAS_VICTIM | INFO | segmentId:{} was VICTIM, so changed to SSD |
| 3179 | ALLOCATOR_SEGMENT_FREED | INFO | segmentId:{} was freed by allocator, free segment count:{} |
| 3180 | VALID_COUNT_UNDERFLOWED | ERROR | segmentId:{} decreasedCount:{} total validCount:{} : UNDERFLOWED |
| 3181 | VALID_COUNT_OVERFLOWED | ERROR | segmentId:{} increasedCount:{} total validCount:{} : OVERFLOWED |
| 3182 | FAILED_TO_ISSUE_ASYNC_METAIO | ERROR | Failed to issue AsyncMetaIo:{} owner:{} |
| 3183 | ERROR_REINIT_WITHOUT_DISPOSE | - | - |
| 3184 | ALLOCATOR_FILE_ERROR | DEBUG | AllocatorCtx file signature is not matched:{} |
| 3185 | GC_STRIPE_ALLOCATED | INFO | Gc stripe is allocated! |
| 3186 | ALLOCATE_GC_VICTIM | INFO | segmentId:{} @AllocateGCVictim, free segment count:{} |
| 3187 | ALLOCATOR_CURRENT_GC_MODE | INFO | Change GC STATE from GCState:{} to URGENT GC MODE, free segment count:{} |
| 3400 | GC_TRIGGERED | - | - |
| 3401 | GC_STARTED | INFO | gc already running |
| 3402 | GC_DONE | INFO | GC done |
| 3403 | GC_CANNOT_START | INFO | cannot start gc |
| 3410 | GC_VICTIM_STRIPE_CONSTRUCTOR | - | - |
| 3411 | GC_LOAD_REVERSE_MAP | INFO | load reversemap for lsid:{} |
| 3412 | GC_LOAD_VALID_BLOCKS | DEBUG | valid blocks loaded, myLsid:{}, cnt:{} |
| 3413 | GC_GET_UNMAP_VSA | INFO | loaded Unmap VSA, volId:{}, rba:{}, stripeId:{}, vsaOffset:{} |
| 3414 | GC_GET_UNMAP_LSA | ERROR | Get Unmap LSA, volId:{}, rba:{}, vsaStripeId:{}, vsaOffset:{}, lsaStripeId:{} |
| 3415 | GC_GET_VICTIM_SEGMENT | DEBUG | trigger start, cnt:{}, victimId:{} |
| 3416 | GC_GET_VALID_BLOCKS | DEBUG | Get valid blocks, (victimStripeId:{}) |
| 3417 | GC_GET_VALID_BLOCK_INFO | - | - |
| 3418 | GC_COPY_COMPLETE | DEBUG | copy complete, id:{} |
| 3419 | GC_GET_BUFFER_FAILED | DEBUG | Get gc buffer failed and retry, victimStripeId:{} |
| 3500 | FC_START | - | - |
| 3501 | FC_TOKEN_DISTRIBUTED | INFO | _DistributeToken userToken:{}, gcToken:{}, userBucket:{}, gcBucket:{} |
| 3502 | FC_TOKEN_OVERFLOW | DEBUG | FlowControl distributed token more than totaltoken totalToken: {}, userToken: {}, gcToken: {} |
| 3503 | FC_NEGATIVE_TOKEN | DEBUG | FlowControl GetToken should be greater than or equal to zero token:{} |
| 3504 | FC_WRONG_STRATEGY | - | - |
| 4000 | MFS_DEBUG_MESSAGE | DEBUG | [MSG ][SubmitIO   ] type={}, req.tagId={}, fd={} |
| 4100 | MFS_INFO_MESSAGE | INFO | {} file has been closed |
| 4200 | MFS_WARNING_INIT_AGAIN | - | catalog is already initialized. |
| 4201 | MFS_SYSTEM_MOUNT_AGAIN | - | - |
| 4202 | MFS_SYSTEM_UNMOUNT_AGAIN | - | - |
| 4203 | MFS_MSG_ENQUEUE_FAILED | - | - |
| 4204 | MFS_COMPACTION_FAILED | - | Compaction couldn't be done due to not enough free space |
| 4300 | MFS_ERROR_MOUNTED | - | - |
| 4301 | MFS_ERROR_UNMOUNTED | - | - |
| 4302 | MFS_ERROR_MESSAGE | - | pread failed... |
| 4303 | MFS_INVALID_PARAMETER | - | MetaIoManager::CheckReqSanity - Invalid OPcode |
| 4304 | MFS_MODULE_NOT_READY | - | - |
| 4305 | MFS_MODULE_ALREADY_READY | - | You attempted to init. volumeMgr duplicately for the given volume type={} |
| 4306 | MFS_MODULE_INIT_FAILED | - | - |
| 4307 | MFS_MODULE_BRINGUP_FAILED | - | - |
| 4308 | MFS_MODULE_NO_MEDIA | - | No registered media info detected. |
| 4309 | MFS_SYSTEM_OPEN_FAILED | - | - |
| 4310 | MFS_INVALID_INFORMATION | - | E2E Check fail!!!! |
| 4311 | MFS_FILE_CREATE_FAILED | ERROR | Map file creation failed, fileName:{} |
| 4312 | MFS_FILE_NOT_FOUND | - | File not found...(given fd={}) |
| 4313 | MFS_FILE_NOT_OPEND | - | Cannot find \'{}\' file in array \'{}\' |
| 4314 | MFS_FILE_OPEN_FAILED | - | File open failed due to an error (e.g.File not found, etc.). rc={} |
| 4315 | MFS_FILE_OPEN_REPETITIONARY | ERROR | You attempt to open fd={} file twice. It is not allowed |
| 4316 | MFS_FILE_CLOSE_FAILED | WARN | VSAMap File close failed, volumeID:{} @VolumeDeleted |
| 4317 | MFS_FILE_FORMAT_FAILED | - | - |
| 4318 | MFS_FILE_DELETE_FAILED | ERROR | MFS File:{} delete failed |
| 4319 | MFS_FILE_ALREADY_LOCKED | - | - |
| 4320 | MFS_FILE_READ_FAILED | - | - |
| 4321 | MFS_FILE_WRITE_FAILED | - | - |
| 4322 | MFS_MEDIA_WROTE_SIZE_NOT_MATCHED | - | Error occurred during AllData erase. |
| 4323 | MFS_MEDIA_SEEK_FAILED | - | lseek failed... |
| 4324 | MFS_MEDIA_READ_FAILED | - | Read failed. return val={}, bytes requested={} |
| 4325 | MFS_MEDIA_WRITE_FAILED | - | write failed... |
| 4326 | MFS_FILE_TRIM_FAILED | - | MFS FILE trim operation has been failed!! |
| 4327 | MFS_MEDIA_MOUNT_FAILED | - | Mount ramdisk failed. Please run as root permission or sudo option... |
| 4328 | MFS_MEDIA_UNMOUNT_FAILED | - | - |
| 4329 | MFS_MEDIA_NULL | - | Mss Disk Place is null |
| 4330 | MFS_META_STORAGE_CREATE_FAILED | - | Failed to create meta storage subsystem |
| 4331 | MFS_META_STORAGE_CLOSE_FAILED | - | - |
| 4332 | MFS_META_STORAGE_NOT_READY | - | - |
| 4333 | MFS_META_VOLUME_CREATE_FAILED | - | Error occurred to create volume (volume id={}) |
| 4334 | MFS_META_VOLUME_OPEN_FAILED | - | Meta volume open is failed. Volume is corrupted or isn't created yet |
| 4335 | MFS_META_VOLUME_CLOSE_FAILED | - | It's failed to close meta volume, arrayName={} |
| 4336 | MFS_META_VOLUME_CLOSE_FAILED_DUE_TO_ACTIVE_FILE | - | Meta filesystem has been unmounted... |
| 4337 | MFS_META_VOLUME_NOT_ENOUGH_SPACE | - | There is no NVRAM and SSD free space |
| 4338 | MFS_META_VOLUME_ALMOST_FULL | DEBUG | NVRAM Volume is almost full |
| 4339 | MFS_META_VOLUME_ALREADY_CLOSED | - | Volume state is not 'Open'. Current state is '{}'. Skip volume close procedure... |
| 4340 | MFS_META_VOLUME_CATALOG_INVALID | - | Volume catalog is invalid. Try to recover broken volume catalog catalogs.. |
| 4341 | MFS_META_SAVE_FAILED | - | Store I/O for MFS MBR content has failed... |
| 4342 | MFS_META_LOAD_FAILED | - | Error occurred while loading filesystem MBR |
| 4343 | MFS_IO_FAILED_DUE_TO_ENQUEUE_FAILED | - | - |
| 4344 | MFS_IO_FAILED_DUE_TO_STOP_STATE | - | - |
| 4345 | MFS_IO_FAILED_DUE_TO_ERROR | - | Failed to enqueue new request... |
| 4346 | MFS_QUEUE_POP_FAILED | - | Item pop failed. rc={} |
| 4347 | MFS_QUEUE_PUSH_FAILED | - | Item push failed. rc={} |
| 4348 | MFS_MAX_FILE_SIZE_NOT_SET | - | - |
| 4349 | MFS_ARRAY_CREATE_FAILED | - | - |
| 4350 | MFS_ARRAY_ADD_FAILED | - | - |
| 4351 | MFS_ARRAY_REMOVE_FAILED | - | - |
| 4352 | MFS_ARRAY_DELETE_FAILED | - | - |
| 4353 | MFS_INVALID_MBR | - | Filesystem MBR has been corrupted or Filesystem cannot be found. |
| 4354 | MFS_NEED_MORE_CONTEXT_SLOT | - | Metafile count={} |
| 4400 | MFS_RECOVERY_CATALOG_FAILED | - | - |
| 4401 | MFS_TEST_FAILED | CRITICAL | Unit test stopped with failure.\n{} {} {} |
| 4500 | VOLUME_EVENT | DEBUG | [Replay] {} block log of volume {} is skipped |
| 4510 | DEVICE_DETACH_EVENT | - | - |
| 4520 | STATE_EVENT | - | - |
| 4580 | LOGGER_FILTER_POLICY_DECODE_FAIL | - | include |
| 4581 | LOGGER_FILTER_POLICY_FILE_NOT_FOUND | - | - |
| 4582 | LOGGER_SET_LEVEL_FAILED | - | - |
| 4600 | QOS_CLI_WRONG_MISSING_PARAMETER | - | Array Name Missing |
| 4601 | QOS_CLI_FE_QOS_DISABLED | - | - |
| 4650 | QOS_SET_EVENT_POLICY | INFO | Rebuild Perf Impact not supported, Set to default Low |
| 4651 | QOS_NOT_SUPPORTED | - | - |
| 5000 | IONVMF_1H | - | RDMA Invalid Private Data Length |
| 5001 | IONVMF_2H | - | RDMA Invalid RECFMT |
| 5002 | IONVMF_3H | - | RDMA Invalid QID |
| 5003 | IONVMF_4H | - | RDMA Invalid HSQSIZE |
| 5004 | IONVMF_5H | - | RDMA Invalid HRQSIZE |
| 5005 | IONVMF_6H | - | RDMA No Resources |
| 5006 | IONVMF_7H | - | RDMA Invalid IRD |
| 5007 | IONVMF_8H | - | RDMA Invalid ORD |
| 5008 | IONVMF_NAMESPACE_ATTACH_FAILED | - | Fail to attach Namespace  |
| 5009 | IONVMF_FAIL_TO_FIND_VOLID | WARN | Fail to parse volume id from bdev name |
| 5010 | IONVMF_FAIL_TO_FIND_ARRAYNAME | WARN | Fail to parse array name from bdev name |
| 5011 | IONVMF_OVERRIDE_UNVMF_IO_HANDLER | WARN | Override unvmf_io_handler |
| 5100 | AFTMGR_CPU_COUNT_NOT_ENOUGH | CRITICAL | Unssatisfied CPU Count |
| 5101 | AFTMGR_DISABLED_CORE | ERROR | Core {} is disabled |
| 5102 | AFTMGR_FAIL_TO_FIND_MASTER_REACTOR | ERROR | Fail to find master reactor |
| 5103 | AFTMGR_FAIL_TO_ALLOCATE_ALL_CPU | - | All core is not assigned for PoseidonOS |
| 5104 | AFTMGR_FAIL_TO_OVERLAP_MASK | CRITICAL | Core mask is overlapped |
| 5105 | AFTMGR_FAIL_TO_PARSING_ERROR | CRITICAL | Cpu allowed list is wrongly set |
| 5106 | AFTMGR_CORE_NOT_SUFFICIENT | ERROR | Cpu is not sufficient |
| 5107 | AFTMGR_NO_EVENT_WORKER_ALLOCATED | ERROR | Cannot find event worker at any NUMA |
| 5108 | AFTMGR_NO_USE_CONFIG | INFO | Use core description from default value |
| 5109 | AFTMGR_USE_CONFIG | INFO | Use core description from config file |
| 5110 | AIO_CONTEXT_NOT_FOUND | - | AIO context is not found |
| 5111 | AIO_FAIL_TO_ALLOCATE_EVENT | - | Event allocation is failed at AIO completion |
| 5112 | AIO_FAIL_TO_ALLOCATE_MEMORY | - | Not sufficient memory at AIO |
| 5113 | AIO_INVALID_AIO_CONTEXT | - | AIO context is null at AIO completion |
| 5114 | AIO_INVALID_AIO_PRIVATE | - | AIOPrivate is null at AIO completion |
| 5115 | AIO_INVALID_BDEV_IO | - | Bdev IO is null at AIO completion |
| 5116 | AIO_INVALID_UBIO | - | Ubio is null at AIO completion |
| 5117 | AIO_IO_FROM_INVALID_THREAD | - | Only reactor can submit IO |
| 5118 | AIO_DEBUG_COMPLETION | - | Aio completion : Rba {}, Size {}, readwrite {}, errortype {} |
| 5119 | AIO_DEBUG_SUBMISSION | - | Aio submission : Rba {}, Size {}, readwrite {} |
| 5120 | AIO_FLUSH_START | INFO_IN_MEMORY | Flush Start in Aio, volume id : {} |
| 5122 | BLKALGN_INVALID_UBIO | ERROR | Block aligning Ubio is null |
| 5123 | BLKHDLR_FAIL_TO_ALLOCATE_EVENT | - | Cannot allocate memory for event |
| 5124 | BLKHDLR_WRONG_IO_DIRECTION | ERROR | Wrong IO direction (only read/write types are suppoered) |
| 5125 | CALLBACK_INVALID_CALLEE | ERROR | Invalid callee for callback |
| 5127 | CALLBACK_TIMEOUT | DEBUG_IN_MEMORY | Callback Timeout. Caller : {} |
| 5128 | CALLBACK_DESTROY_WITHOUT_EXECUTED | WARN | Callback destroy without executed : {} |
| 5129 | CALLBACK_DESTROY_WITH_ERROR | WARN | Callback Error : {} |
| 5130 | EVENTFRAMEWORK_FAIL_TO_ALLOCATE_EVENT | WARN | Fail to allocate spdk event |
| 5131 | EVENTFRAMEWORK_INVALID_EVENT | ERROR | Invalid Event to send |
| 5132 | EVENTFRAMEWORK_INVALID_REACTOR | ERROR | Reactor {} is not processable |
| 5134 | EVTARG_WRONG_ARGUMENT_ACCESS | - | Request for getting wrong argument index |
| 5135 | EVTARG_WRONG_ARGUMENT_ADD | - | Request for adding wrong argument index |
| 5136 | EVTQ_INVALID_EVENT | - | Input event is null at EventQueue |
| 5137 | EVTSCHDLR_INVALID_WORKER_ID | - | WorkerID is not valid at EventScheduler |
| 5138 | FLUSHREAD_DEBUG_SUBMIT | - | Flush Read Submission StartLSA.stripeId : {} blocksInStripe : {} |
| 5139 | FLUSHREAD_FAIL_TO_ALLOCATE_MEMORY | ERROR | Fail to allocate memory |
| 5140 | FLUSH_DEBUG_SUBMIT | DEBUG_IN_MEMORY | translator in Flush Submission has error code : {} stripeId : {} |
| 5141 | FLUSH_DEBUG_COMPLETION | DEBUG_IN_MEMORY | Flush Completion vsid : {} userstripeid : {} userArea : {} |
| 5142 | FREEBUFPOOL_FAIL_TO_ALLOCATE_MEMORY | ERROR | Fail to allocate memory |
| 5143 | IOAT_CONFIG_INVALID | - | Invalid ioat count for numa |
| 5144 | IOATAPI_FAIL_TO_FINALIZE | ERROR | Fail to finalize IOAT |
| 5145 | IOATAPI_FAIL_TO_INITIALIZE | ERROR | Fail to initialize IOAT |
| 5146 | IOATAPI_FAIL_TO_SUBMIT_COPY | - | Fail to request IOAT copy to reactor |
| 5147 | IOATAPI_DISABLED | INFO | IOAT disabled |
| 5148 | IOATAPI_ENABLED | - | IOAT enabled |
| 5149 | IODISPATCHER_INVALID_CPU_INDEX | ERROR | Invalid CPU index {} |
| 5150 | IODISPATCHER_INVALID_PARM | ERROR | Invalid Param in submit IO |
| 5151 | MERGER_INVALID_SPLIT_REQUESTED | - | Requested Ubio index exceeds Ubio count of Merger |
| 5152 | RBAMGR_WRONG_VOLUME_ID | - | Volume ID is not valid at RBAStateManager |
| 5153 | RDCMP_INVALID_ORIGIN_UBIO | - | Origin Ubio is null at ReadCompleteteHandler |
| 5154 | RDCMP_INVALID_UBIO | ERROR | Ubio is null at ReadCompleteHandler |
| 5155 | RDCMP_READ_FAIL | ERROR | Uncorrectable data error |
| 5156 | RDHDLR_INVALID_UBIO | - | Read Ubio is null |
| 5157 | RDHDLR_READ_DURING_REBUILD | - | Data is read from which is under rebuild process |
| 5158 | SCHEDAPI_COMPLETION_POLLING_FAIL | ERROR | Fail to poll ibof completion |
| 5159 | SCHEDAPI_NULL_COMMAND | ERROR | Command from bdev is empty |
| 5160 | SCHEDAPI_SUBMISSION_FAIL | ERROR | Fail to submit ibof IO |
| 5161 | SCHEDAPI_WRONG_BUFFER | ERROR | Single IO command should have a continuous buffer |
| 5162 | STRIPE_INVALID_VOLUME_ID | - | VolumeId is invalid |
| 5163 | REF_COUNT_RAISE_FAIL | - | When Io Submit, refcount raise fail |
| 5164 | TRANSLATE_CONVERT_FAIL | ERROR | Translate() or Convert() is failed |
| 5165 | TRANSLATE_CONVERT_FAIL_IN_SYSTEM_STOP | - | Translate() or Convert() is failed in system stop |
| 5166 | TRSLTR_INVALID_BLOCK_INDEX | ERROR | Block index exceeds block count at Translator |
| 5167 | TRSLTR_WRONG_ACCESS | ERROR | Only valid for single block Translator |
| 5168 | TRSLTR_WRONG_VOLUME_ID | ERROR | Volume ID is not valid at Translator |
| 5169 | VOLUMEIO_DEBUG_SUBMIT | - | Volume IO Submit, Rba : {}, Size : {}, bypassIOWorker : {} |
| 5170 | UBIO_DEBUG_CHECK_VALID | - | Ubio Check before Submit, Rba : {}, Pba.lba : {}, Size : {}, Ublock : {}, ReferenceIncreased :{} sync :{} Direction : {} |
| 5171 | UBIO_DEBUG_COMPLETE | - | Ubio Complete Pba.lba : {} Size : {} Ublock : {} ErrorType : {} Direction : {} |
| 5172 | UBIO_ALREADY_SYNC_DONE | - | Mark done to already completed Ubio |
| 5173 | UBIO_CALLBACK_EVENT_ALREADY_SET | - | Callback Event is already set for Ubio |
| 5174 | UBIO_DUPLICATE_IO_ABSTRACTION | - | Double set request context |
| 5175 | UBIO_FAIL_TO_ALLOCATE_MEMORY | - | Not sufficient memory at Ubio |
| 5176 | UBIO_FREE_UNALLOWED_BUFFER | ERROR | Cannot free unallowed data buffer |
| 5177 | UBIO_INVALID_GC_PROGRESS | - | Invalid GC progress |
| 5178 | UBIO_INVALID_POS_IO | - | Invalid ibof IO |
| 5179 | UBIO_INVALID_IO_STATE | - | Invalid IO state |
| 5180 | UBIO_INVALID_LSID | ERROR | Invalid LSID for Ubio |
| 5181 | UBIO_INVALID_ORIGIN_UBIO | ERROR | Invalid original Ubio |
| 5182 | UBIO_INVALID_ORIGINAL_CORE | - | Invalid origin core for Ubio |
| 5183 | UBIO_INVALID_PBA | ERROR | Invalid PBA for Ubio |
| 5184 | UBIO_INVALID_RBA | ERROR | Invalid RBA for Ubio |
| 5185 | UBIO_INVALID_UBIO_HANDLER | - | Invalid ubio handler |
| 5186 | UBIO_INVALID_VOLUME_ID | ERROR | Invalid volume ID for Ubio |
| 5187 | UBIO_INVALID_VSA | ERROR | Invalid VSA for Ubio |
| 5188 | UBIO_INVALID_ARRAY_NAME | - | Invalid Array Name for Ubio |
| 5189 | UBIO_INVALID_DEVICE | ERROR | Invalid UblockDevice for Ubio |
| 5190 | UBIO_NO_COMPLETION_FOR_FRONT_END_EVENT | - | No explicit implementation for front-end complete type |
| 5191 | UBIO_REMAINING_COUNT_ERROR | - | Completion count exceeds remaining count |
| 5192 | UBIO_REQUEST_NULL_BUFFER | ERROR | Requested buffer of Ubio is null |
| 5193 | UBIO_REQUEST_OUT_RANGE | ERROR | Requested buffer of Ubio is out of range |
| 5194 | UBIO_SYNC_FLAG_NOT_SET | - | Mark done is only valid for sync Ubio |
| 5195 | UBIO_WRONG_SPLIT_SIZE | ERROR | Invalid size of Ubio split request |
| 5196 | URAM_BACKUP_FILE_NOT_EXISTS | - | URAM backup file doesn't exist, ignoring backup restoration. |
| 5197 | URAM_BACKUP_FILE_OPEN_FAILED | - | URAM backup file does exist but opening is failed |
| 5198 | URAM_BACKUP_FILE_READ_FAILED | - | Read page #  |
| 5199 | URAM_BACKUP_FILE_STAT_FAILED | - | Cannot get URAM backup file stat |
| 5200 | URAM_CONFIG_FILE_OPEN_FAILED | ERROR | Cannot open uram config file |
| 5201 | URAM_COMPLETION_TIMEOUT | - | URAM completion checking timed out |
| 5202 | URAM_FAIL_TO_CLOSE | - | Fail to close URAM |
| 5203 | URAM_FAIL_TO_OPEN | - | Fail to open URAM |
| 5204 | URAM_FAIL_TO_REQUEST_IO | - | Fail to request nvram IO to reactor |
| 5205 | URAM_FAIL_TO_RETRY_IO | - | Fail to retry nvram IO to reactor |
| 5206 | URAM_PENDING_IO_NOT_FOUND | - | Pending Ubio not found: DeviceContext: {}, Ubio: {} |
| 5207 | URAM_RESTORING_FAILED | - | Cannot open URAM driver for restoring |
| 5208 | URAM_RESTORING_PAGE_DONE | INFO | Restoring Hugepage #{} done. |
| 5209 | URAM_SUBMISSION_FAILED | - | IO submission failed |
| 5210 | URAM_SUBMISSION_TIMEOUT | - | Could not submit the given IO request in time |
| 5211 | WRCMP_FAIL_TO_ALLOCATE_MEMORY | - | Fail to allocate memory |
| 5212 | WRCMP_INVALID_SPLIT_UBIO | - | Split Ubio is null at WriteCompleting state |
| 5213 | WRCMP_INVALID_STRIPE | ERROR | Stripe is null at WriteCompleting state |
| 5214 | WRCMP_IO_ERROR | ERROR | Write is failed at WriteCompleting state |
| 5215 | WRCMP_WRITE_WRAPUP_FAILED | - | Write wrapup failed at WriteCompleting state |
| 5216 | WRCMP_MAP_UPDATE_FAILED | ERROR | Write wraup failed at map update |
| 5217 | IOCONTROL_REBUILD_FAIL | - | Rebuild Read Failed |
| 5218 | WRHDLR_ALLOC_WRITE_BUFFER_FAILED | - | Not enough room for write buffer, this write is going to be rescheduled |
| 5219 | WRHDLR_FAIL_TO_ALLOCATE_MEMORY | - | Not sufficient memory |
| 5220 | WRHDLR_FAIL_BY_SYSTEM_STOP | ERROR | System Stop incurs write fail |
| 5221 | WRHDLR_INVALID_UBIO | - | Write Ubio is null |
| 5222 | WRHDLR_NO_FREE_SPACE | DEBUG | No free space in write buffer |
| 5223 | WRHDLR_DEBUG_READ_OLD | - | Read Old Block in Write Path Rba : {} Size : {} Vsa.id : {} Vsa.offset :{} Lsid.id : {}, Lsid.Loc :{} Remained : {} |
| 5224 | WRHDLR_DEBUG_READ_OLD_COMPLETE | - | Read Old Block in Write Path Rba : {} Size : {} Lsid.id : {} alignOffset : {} alignSize : {} |
| 5225 | WRWRAPUP_EVENT_ALLOC_FAILED | ERROR | Flush Event allocation failed at WriteWrapup state |
| 5226 | WRWRAPUP_STRIPE_NOT_FOUND | ERROR | Stripe #{} not found at WriteWrapup state |
| 5227 | WRWRAPUP_DEBUG_STRIPE | - | Write Completion, Vsa : {}, remain : {} stripe vid : {} |
| 5228 | BLKMAP_DEBUG_UPDATE_REQUEST | - | Block Map Update Request : StartRba : {} Vsa : {} blCnt : {} volumeId :{}, isGC :{} isOldData :{} |
| 5229 | BLKMAP_DEBUG_UPDATE | - | Block Map Update Request : StartRba : {} Vsa : {} blCnt : {} volumeId :{}, isGC :{} isOldData :{} |
| 5230 | FLUSH_HANDLING_ENABLED | INFO | Flush command handling is enabled |
| 5231 | FLUSH_HANDLING_DISABLED | INFO | Flush command handling is disabled as journal is enabled |
| 5232 | FLUSH_CMD_MAPPER_FLUSH_FAILED | ERROR | Failed to start flushing dirty vsamap pages |
| 5233 | FLUSH_CMD_ALLOCATOR_FLUSH_FAILED | ERROR | Failed to start flushing allocator meta pages. Error code {} |
| 5234 | FLUSH_CMD_ONGOING | DEBUG | Flush command started on volume {} |
| 5300 | IOWORKER_OPERATION_NOT_SUPPORTED | - | Requested operation is not supported by IOWorker: {} |
| 5301 | IOWORKER_DEVICE_ADDED | INFO | {} has been added to IOWorker{} |
| 5302 | IOWORKER_UNDERFLOW_HAPPENED | ERROR | Command completed more than submitted: current outstanding: {}, completion count: {} |
| 5303 | IOSMHDLR_BUFFER_NOT_ENOUGH | - | Buffer is not enough to proceed with IO request |
| 5304 | IOSMHDLR_BUFFER_NOT_ALIGNED | - | Buffer size is not properly aligned to proceed with IO request |
| 5305 | IOSMHDLR_OPERATION_NOT_SUPPORTED | - | Requested operation is not supported by IOSubmitHandler |
| 5306 | IOSMHDLR_DEBUG_ASYNC_READ | - | IOSubmitHandler Async Read : physical lba : {}  blkcnt :{} bufferIndex : {} partitionToIO : {} ret :{} |
| 5307 | IOSMHDLR_COUNT_DIFFERENT | WARN | IOSubmitHandler Async BufferCounts are different :  total of each entries {}  blkcnt :{} |
| 5308 | IOSMHDLR_DEBUG_ASYNC_WRITE | - | IOSubmitHandler Async Write : physical lba : {} blkcnt : {} partitionToIO : {} totalcnt :{} |
| 5309 | IOSMHDLR_ARRAY_LOCK | - | IOSubmitHandler Array Locking type : {} stripeId :{} count :{} |
| 5310 | IOSMHDLR_ARRAY_UNLOCK | - | IOSubmitHandler Array Unlocking type : {} stripeId :{} |
| 5311 | I_IOSMHDLR_NULLPTR | WARN | interface of io submit handler already registered. |
| 5312 | DEVICE_OPEN_FAILED | - | Device open failed: Name: {}, Desc: {}, libaioContextID: {}, events: {} |
| 5313 | DEVICE_CLOSE_FAILED | - | Device close failed: Name: {}, Desc: {}, libaioContextID: {}, events: {} |
| 5314 | DEVICE_SCAN_FAILED | - | Error occurred while scanning newly detected device: {} |
| 5315 | DEVICE_SCAN_IGNORED | - | Device: {} ignored while scanning since the size: {} Bytes is lesser than required: > {} Bytes |
| 5316 | DEVICE_SUBMISSION_FAILED | - | IO submission failed: Device: {}, Op: {}, Start LBA: {}, Size in Bytes: {} |
| 5317 | DEVICE_SUBMISSION_QUEUE_FULL | - | IO submission pending since the queue is full: Device: {} |
| 5318 | DEVICE_SUBMISSION_TIMEOUT | - | Could not submit the given IO request in time to device: {}, Op: {}, Start LBA: {}, Size in Bytes: {} |
| 5319 | DEVICE_COMPLETION_FAILED | ERROR | Error: {} occurred while getting IO completion events from device: {} |
| 5320 | DEVICE_THREAD_REGISTERED_FAILED | ERROR | Error: {} register device context failed: {} |
| 5321 | DEVICE_THREAD_UNREGISTERED_FAILED | - | Error: {} unregister device context failed: {} |
| 5322 | DEVICE_OPERATION_NOT_SUPPORTED | - | Requested operation: {} is not supported by DeviceDriver |
| 5323 | DEVICE_PENDING_IO_NOT_FOUND | - | Pending Ubio not found: DeviceContext: {}, Ubio: {} |
| 5325 | DEVICE_OVERLAPPED_SET_IOWORKER | WARN | Overlapped setting for ioworker for single device: {}  |
| 5326 | DEVICE_NULLPTR_IOWORKER | WARN | Overlapped setting for ioworker for single device: {}  |
| 5327 | QOSMGR_REGISTER_POLLER_FAILED | ERROR | Failed to register Qos poller on reactor #: {} |
| 5328 | QOSMGR_UNREGISTER_POLLER_FAILED | ERROR | Failed to un-register Qos poller on reactor #: {} |
| 5329 | UNVME_SSD_DEBUG_CREATED | DEBUG | Create Ublock, Pointer : {} |
| 5330 | UNVME_SSD_SCAN_FAILED | ERROR | Failed to Scan uNVMe devices |
| 5331 | UNVME_SSD_SCANNED | INFO | uNVMe Device has been scanned: {}, {} |
| 5332 | UNVME_SSD_ATTACH_NOTIFICATION_FAILED | ERROR | Failed to notify uNVMe device attachment: Device name: {} |
| 5333 | UNVME_SSD_DETACH_NOTIFICATION_FAILED | ERROR | Failed to notify uNVMe device detachment: Device name: {} |
| 5334 | UNVME_SSD_OPEN_FAILED | ERROR | uNVMe Device open failed: namespace #{} |
| 5335 | UNVME_SSD_CLOSE_FAILED | ERROR | uNVMe Device close failed: namespace #{} |
| 5336 | UNVME_SSD_UNDERFLOW_HAPPENED | ERROR | Command completed more than submitted: current outstanding: {}, completion count: {} |
| 5337 | UNVME_SUBMISSION_FAILED | ERROR | Command submission failed: namespace: {}, errorCode: {} |
| 5338 | UNVME_SUBMISSION_QUEUE_FULL | - | IO submission pending since the queue is full |
| 5339 | UNVME_SUBMISSION_RETRY_EXCEED | - | Could not submit the given IO request within limited count : lba : {} sectorCount : {} deviceName : {} namespace {} |
| 5340 | UNVME_COMPLETION_TIMEOUT | WARN | uNVMe completion checking timed out: SN: {} |
| 5341 | UNVME_COMPLETION_FAILED | ERROR | Command completion error occurred: namespace: {}, errorCode: {} |
| 5342 | UNVME_OPERATION_NOT_SUPPORTED | WARN | Requested operation: {} is not supported by uNVMe DeviceDriver |
| 5343 | UNVME_CONTROLLER_FATAL_STATUS | WARN | Controller Fatal Status, reset required: SN: {} |
| 5344 | UNVME_CONTROLLER_RESET_FAILED | - | Controller Reset Failed: SN: {} |
| 5345 | UNVME_CONTROLLER_RESET | ERROR | Controller Reset Failed: SN: {} |
| 5346 | UNVME_SUBMITTING_CMD_ABORT | INFO | Requesting command abort: SN: {}, QPair: {}, CID: {} |
| 5347 | UNVME_ABORT_TIMEOUT | INFO | Abort Also Timeout: SN: {}, QPair: {}, CID: {} |
| 5348 | UNVME_ABORT_SUBMISSION_FAILED | - | Failed to submit command abort: SN: {}, QPair: {}, CID: {} |
| 5349 | UNVME_ABORT_COMPLETION_FAILED | ERROR | Failed to complete command abort: SN: {} |
| 5350 | UNVME_ABORT_COMPLETION | ERROR | Failed to complete command abort: SN: {} |
| 5351 | UNVME_DO_NOTHING_ON_TIMEOUT | - | Do nothing on command timeout: SN: {} |
| 5352 | UNVME_ABORT_TIMEOUT_FAILED | - | Requesting command abort: SN: {}, QPair: {}, CID: {} |
| 5353 | UNVME_NOT_SUPPORTED_DEVICE | WARN | Device {} is not supported. Sector size is not {} |
| 5354 | UNVME_DEBUG_RETRY_IO | WARN | Retry IO in unvme_drv, startLBA: {}, sectorCount : {}, direction : {}, deviceName : {} |
| 5355 | UNVME_DEBUG_REQUEST_IO | DEBUG_IN_MEMORY | Request IO in unvme_drv, startLBA: {}, sectorCount : {}, direction : {} deviceName : {} ret : {} |
| 5356 | UNVME_DEBUG_COMPLETE_IO | DEBUG_IN_MEMORY | Complete IO in unvme_drv, startLBA: {}, sectorCount : {}, direction : {}, sc : {}, sct : {} deviceName : {} |
| 5357 | DEVICE_CONTEXT_NOT_FOUND | DEBUG | Device context is not found |
| 5358 | DEVCTX_ALLOC_TIMEOUT_CHECKER_FAILED | - | Allocating TimeoutChecker for pending error list failed. |
| 5359 | IOCTX_ERROR_KEY_NOT_SET | ERROR | Key for pending ERROR was not set, before! |
| 5360 | IOCTX_IO_KEY_NOT_SET | ERROR | Key for pending IO was not set, before! |
| 5361 | IOQ_ENQUEUE_NULL_UBIO | WARN | Enqueue null ubio |
| 5362 | BUFFER_ENTRY_OUT_OF_RANGE | ERROR | Block / Chunk Index excceds block count of BufferEntry |
| 5363 | NFLSH_ERROR_DETECT | ERROR | Failed to proceed stripe map update request event: {} |
| 5364 | NFLSH_EVENT_ALLOCATION_FAILED | ERROR | FlushCompletion for vsid:  |
| 5365 | NFLSH_EVENT_MAP_UPDATE_FAILED | ERROR | FlushCompletion for vsid:  |
| 5366 | NFLSH_STRIPE_NOT_IN_WRITE_BUFFER | ERROR | Stripe #{} is not in WriteBuffer. |
| 5367 | NFLSH_STRIPE_DEBUG | DEBUG_IN_MEMORY | Stripe Map Update Request : stripe.vsid : {} writeBufferArea : {} wbStripeid : {} |
| 5368 | NFLSH_STRIPE_DEBUG_UPDATE | ERROR | Stripe Map Update Request : stripe.vsid : {} logWriteRequestSuccess : {} |
| 5369 | FLUSH_WRAPUP_STRIPE_NOT_IN_USER_AREA | ERROR | Stripe #{} is not in UserArea. |
| 5370 | STRIPEPUTEVT_STRIPE_NOT_IN_NORMAL_POOL | - | Stripe #{} is not in NormalStripePool. |
| 5500 | UNVME_DAEMON_START | DEBUG | spdk daemon started |
| 5501 | UNVME_DAEMON_FINISH | DEBUG | spdk daemon stopped |
| 5502 | UNVME_REGISTER_CTRL | ERROR | ctrlr_entry malloc |
| 5503 | UNVME_REGISTER_NS | INFO | Controller - MDTS: {} |
| 5504 | UNVME_PROBE_CALLBACK | INFO | Probing {}  |
| 5505 | UNVME_INIT_SCAN_CALLBACK | INFO | Attaching to {} |
| 5506 | UNVME_ATTACH_CALLBACK | WARN | spdk device attachment detected |
| 5507 | UNVME_DETACH_CALLBACK | WARN | spdk device detachment detected |
| 5508 | UNVME_SPDK_DETACH | WARN | SpdkDetach - ns == NULL |
| 5509 | UNVME_CLEAN_UP | INFO | Detaching SPDK controlllers! |
| 5510 | UNVME_INIT_CONTROLLER | ERROR | spdk_nvme_probe() failed |
| 5511 | UNVME_MAX_RETRY_EXCEED | WARN | Backend's Retry Count from Configuration File Excceds Max Retry Count : {}, Input : {} |
| 5512 | UNVME_MAX_TIMEOUT_EXCEED | WARN | SSD Timeout usec From Configuration File Excceds Max timeout : {}, Input : {} |
| 5513 | DEVICEMGR_ATTACH | INFO | DeviceAttachEventHandler: {} |
| 5514 | DEVICEMGR_CLEAR_DEVICE | INFO | devices has been cleared sucessfully |
| 5515 | DEVICEMGR_START_MONITOR | INFO | Start Monitoring |
| 5516 | DEVICEMGR_STOP_MONITOR | INFO | StopMonitoring |
| 5517 | DEVICEMGR_CHK_DUPLICATE | WARN | DEVICE DUPLICATED Name:{}, SN: {} |
| 5518 | DEVICEMGR_INIT_SCAN | INFO | InitScan begin |
| 5519 | DEVICEMGR_REMOVE_DEV | ERROR | device not found |
| 5520 | DEVICEMGR_RESCAN | INFO | ReScan begin |
| 5521 | DEVICEMGR_GETDEV | DEBUG | Device Found: {} |
| 5522 | DEVICEMGR_LISTDEV | INFO | ListDevs, count: {} |
| 5523 | DEVICEMGR_DETACH | WARN | DetachDevice - unknown device or already detached: {} |
| 5524 | DEVICEMGR_SETUPMODEL | DEBUG | _SetupThreadModel |
| 5525 | DEVICEMGR_DEVICE_NOT_FOUND | DEBUG | Device might be not scanned, continue to exit |
| 5700 | RESOURCE_MANAGER_DEBUG_MSG | DEBUG | Failed to return buffer. Buffer is Null |
| 8500 | INVALID_BUFFER_IN_SMART_LOG | ERROR | Invalid buffer in smart log page |
| 8501 | NO_DISK_IN_ARRAY | ERROR | No Device in Array |
| 9000 | SYSTEM_SPACE_INFO | - | - |
| 9001 | SYSTEM_VERSION | INFO | POS Version {} |
| 10000 | DEBUG_MEMORY_CHECK_DOUBLE_FREE | ERROR | double free {} |
| 10001 | DEBUG_MEMORY_CHECK_INVALID_ACCESS | ERROR | invalid access base : {} size : {} |
| 2147483646 | EVENT_ID_MAPPING_WRONG | ERROR | Mapping EventID with Message is wrong: Requested eventId: {},  |
| 2147483647 | RESERVED | ERROR | Reserved, Not defined, yet!! |
