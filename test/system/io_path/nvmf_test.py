#!/usr/bin/env python3

import os
import nvmf_common
import common_test_lib
import argparse

#######################################################################
ibof_root = os.path.dirname(os.path.abspath(__file__)) + "/../../../"
spdk_path = ibof_root + "lib/spdk/"
spdk_rpc_path = spdk_path + "scripts/rpc.py"
default_log_path = "pos.log"
default_print_on = False
default_transport = "tcp" #or rdma
default_target_ip = "10.100.11.16"

#######################################################################

def create_transport_test(EXPECT, trtype, b, n):
    nvmf = nvmf_common.Nvmf("[ Create Transport ", EXPECT, spdk_rpc_path)
    common_test_lib.start_pos(args.log_path, ibof_root)

    ret = nvmf.create_transport(trtype, b, n)

    common_test_lib.expect_true(ret, "Transport Creation")
    common_test_lib.terminate_pos(ibof_root, stdout_type)
    success = (EXPECT==ret)
    common_test_lib.print_result("[ Create Transport Test ]", success)
    return success

def create_subsystem_test(EXPECT, nqn, serial_num, model_num, max_ns,
        allow_any_host):
    nvmf = nvmf_common.Nvmf("[ Create Subsystem ", EXPECT, spdk_rpc_path)
    common_test_lib.start_pos(args.log_path, ibof_root)

    ret = nvmf.create_subsystem(nqn, serial_num, model_num, max_ns,
            allow_any_host)

    common_test_lib.expect_true(ret, "Subsystem Creation")
    common_test_lib.terminate_pos(ibof_root, stdout_type)
    success = (EXPECT == ret)
    common_test_lib.print_result("[ Create Subsystem Test ]", success)
    return success

def delete_subsystem_test(EXPECT, nqn, allow_any_host):
    nvmf = nvmf_common.Nvmf("[ Delete Subsystem ", EXPECT, spdk_rpc_path)
    common_test_lib.start_pos(args.log_path, ibof_root)

    ret = nvmf.create_subsystem("nqn.2019-04.pos:subsystem1",\
            "POS00000000000001","POS_VOLUME_EXTENTION", "256", allow_any_host)
    common_test_lib.expect_true(ret, "Subsystem Creation")
    
    ret = nvmf.delete_subsystem(nqn)

    common_test_lib.expect_true(ret, "Subsystem Deletion")
    common_test_lib.terminate_pos(ibof_root, stdout_type)
    success = (EXPECT == ret)
    common_test_lib.print_result("[ Delete Subsystem Test ]", success)
    return success

def add_subsystem_listener_test(EXPECT, nqn, trtype, traddr, trsvid, allow_any_host):
    nvmf = nvmf_common.Nvmf("[ Add Subsystem Listener ", EXPECT, spdk_rpc_path)
    common_test_lib.start_pos(args.log_path, ibof_root)

    ret = nvmf.create_transport(trtype, "64", "2048")
    common_test_lib.expect_true(ret, "Transport Creation")

    ret = nvmf.create_subsystem("nqn.2019-04.pos:subsystem1",\
            "POS00000000000001","POS_VOLUME_EXTENTION", "256", allow_any_host)
    common_test_lib.expect_true(ret, "Subsystem Creation")

    ret = nvmf.add_subsystem_listener(nqn, trtype, traddr, trsvid)
    
    common_test_lib.expect_true(ret, "Subsystem Listener Added")
    common_test_lib.terminate_pos(ibof_root, stdout_type)
    success = (EXPECT == ret)
    common_test_lib.print_result("[ Add Subsystem Listener Test ]", success)
    return success

def subsystem_namespace_mapping_test(trtype, traddr, subsystem_cnt, volume_cnt):
    nvmf = nvmf_common.Nvmf("[ Subsystem-Namespace Mapping ", True, spdk_rpc_path)
    
    bringup_argument = {
        'log_path' : args.log_path,
        'ibof_root' : ibof_root,
        'transport' : trtype,
        'target_ip' : traddr,
        'subsystem_cnt' : subsystem_cnt,
        'volume_cnt' : volume_cnt,
        'stdout_type' : stdout_type
        }

    common_test_lib.bringup_multiple_volume(**bringup_argument)

    map = nvmf.get_subsystem_ns_map()
    ret = nvmf.check_subsystem_ns_mapping(map, "subsystem1", "bdev_3")
    success = (ret == True)
    ret = nvmf.check_subsystem_ns_mapping(map, "subsystem3", "bdev_3")
    success &= (ret == False)

    common_test_lib.terminate_pos(ibof_root, stdout_type)
    common_test_lib.print_result("[ Subsystem-Namespace Mapping Test ]", success)
    return success

def parse_arguments():
    parser = argparse.ArgumentParser(description='Test NVMf functions')
    parser.add_argument('-f', '--fabric_ip', default=default_target_ip,\
            help='Set target IP, default: ' + default_target_ip)
    parser.add_argument('-l', '--log_path', default=default_log_path,\
            help='Set path for log file, default: ' + default_log_path)
    parser.add_argument('-t', '--transport', default=default_transport,
            help='Set transport, default: ' + default_transport)
    parser.add_argument('-p', '--print_log', default=default_print_on,
            help='Set printing log or not, default: ' + str(default_print_on))
    global args
    args = parser.parse_args()
    print (args)

if __name__ == "__main__":
    parse_arguments()
    common_test_lib.clear_env()
    stdout_type = common_test_lib.set_stdout_type(args.print_log)

    success = create_transport_test(True, args.transport, "64", "2048")
    success &= create_transport_test(False, "wrong_transport_type", "64",
            "2048")
    allow_any_host = True
    success &= create_subsystem_test(True, "nqn.2019-04.pos:subsystem1",\
            "POS00000000000001","POS_VOLUME_EXTENTION", "256", allow_any_host)
    success &= create_subsystem_test(False, "nqn.2019-04",\
            "POS00000000000001","POS_VOLUME_EXTENTION", "256", allow_any_host)

    success &= delete_subsystem_test(True, "nqn.2019-04.pos:subsystem1",
                    allow_any_host)
    success &= delete_subsystem_test(False, "nqn.2019", allow_any_host)

    success &= add_subsystem_listener_test(True, "nqn.2019-04.pos:subsystem1",
            args.transport, args.fabric_ip, "1158", allow_any_host)
    success &= add_subsystem_listener_test(False, "nqn.2019-04.pos:subsystem2",
            args.transport, args.fabric_ip, "1158", allow_any_host)
    
    success &= subsystem_namespace_mapping_test(args.transport, args.fabric_ip,
            "3", "5")
    if success:
        exit(0)
    else:
        exit(-1)
