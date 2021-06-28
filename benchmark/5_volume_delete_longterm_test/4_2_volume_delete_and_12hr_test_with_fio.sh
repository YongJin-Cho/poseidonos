#!/bin/bash
source ../config.sh

root_dir=${ROOT_DIR}
target_ip_1=${TARGET_IP_1}
target_ip_2=${TARGET_IP_2}
target_nic_1=${TARGET_NIC_1}
target_nic_2=${TARGET_NIC_2}
volume_cnt=${VOLUME_CNT}
volume_gb_size=${VOLUME_GB_SIZE}
volume_byte_size=${VOLUME_BYTE_SIZE}
# second
write_fill_seq_io_time=${WRITE_FILL_SEQ_IO_TIME}
write_fill_rand_io_time=${WRITE_FILL_RAND_IO_TIME}
seq_io_time=${VOL_DEL_SEQ_IO_TIME}
rand_io_time=${VOL_DEL_RAND_IO_TIME}
delete_volume_cnt=${DELETE_VOLUME_CNT}
remain_volume_cnt=${REMAIN_VOLUME_CNT}
deleted_start_volume_num=$(($remain_volume_cnt+1))

init1_id=${INIT_1_ID}
init2_id=${INIT_2_ID}
init1_pw=${INIT_1_PW}
init2_pw=${INIT_2_PW}
init1_ip=${INIT_1_IP}
init2_ip=${INIT_2_IP}

init1_fio_conf_dir=${INIT1_FIO_CONF_DIR}
init2_fio_conf_dir=${INIT2_FIO_CONF_DIR}
init1_fio_engine=${INIT1_FIO_ENGINE}
init2_fio_engine=${INIT2_FIO_ENGINE}
init1_files=${INIT1_FILES}
init2_files=${INIT2_FILES}


volume_delete()
{
    for ((i=${deleted_start_volume_num};i<=${volume_cnt};i++))
    do
        $root_dir/bin/cli volume unmount --name vol$i --array POSArray
        $root_dir/bin/cli volume delete --name vol$i --array POSArray
    done
}

volume_create_and_mount()
{
for ((i=${deleted_start_volume_num};i<=${volume_cnt};i++))
do
    sudo $ROOT_DIR/bin/cli volume create --name vol$i --size $volume_byte_size --maxiops 0 --maxbw 0 --array POSArray
    sudo $ROOT_DIR/bin/cli volume mount --name vol$i --array POSArray
done
}

write_fill_pos_before_vol_delete()
{
sudo ./create_test_config.sh -a ${target_ip_1} -b ${target_ip_2} -v ${volume_cnt} -S ${volume_gb_size} -s ${write_fill_seq_io_time} -r ${write_fill_rand_io_time} -f ${init1_fio_conf_dir} -e ${init1_fio_engine}
sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "mkdir -p ${init1_fio_conf_dir}"
for file in ${init1_files}
do
    echo ${file}
    sshpass -p ${init1_pw} scp -o StrictHostKeyChecking=no ${file} ${init1_id}@${init1_ip}:${init1_fio_conf_dir}
done
echo "prepare init 1 finish"

echo "test start"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "mkdir -p ${init2_fio_conf_dir}"
for file in ${init2_files}
do
    echo ${file}
    sshpass -p ${init2_pw} scp -o StrictHostKeyChecking=no ${file} ${init2_id}@${init2_ip}:${init2_fio_conf_dir}
done
echo "prepare init 2 finish"

if [ ${write_fill_seq_io_time} -gt 1 ]
then
echo "sequential write start"
sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "cd ${init1_fio_conf_dir}; echo ${init1_pw} | sudo -S nohup fio ./sw_tcp_init1.conf > /dev/null 2>&1 &"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "cd ${init2_fio_conf_dir}; echo ${init2_pw} | sudo -S nohup fio ./sw_tcp_init2.conf > /dev/null 2>&1"
echo "sequential write sleep"
sleep 1
fi

if [ ${write_fill_rand_io_time} -gt 1 ]
then
echo "random write start"
sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "cd ${init1_fio_conf_dir}; echo ${init1_pw} | sudo -S nohup fio ./rw_tcp_init1.conf > /dev/null 2>&1 &"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "cd ${init2_fio_conf_dir}; echo ${init2_pw} | sudo -S nohup fio ./rw_tcp_init2.conf > /dev/null 2>&1"
echo "random write sleep"
fi

sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "echo ${init1_pw} | sudo -S rm ${init1_fio_conf_dir}/*.log"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "echo ${init2_pw} | sudo -S rm ${init2_fio_conf_dir}/*.log"
}

longterm_test_after_vol_delete()
{
volume_delete

sudo ./create_test_config.sh -a ${target_ip_1} -b ${target_ip_2} -v ${remain_volume_cnt} -S ${volume_gb_size} -s ${seq_io_time} -r ${rand_io_time} -f ${init1_fio_conf_dir} -e ${init1_fio_engine}

sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "mkdir -p ${init1_fio_conf_dir}"
for file in ${init1_files}
do
    echo ${file}
    sshpass -p ${init1_pw} scp -o StrictHostKeyChecking=no ${file} ${init1_id}@${init1_ip}:${init1_fio_conf_dir}
done
echo "prepare init 1 finish"

echo "test start"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "mkdir -p ${init2_fio_conf_dir}"
for file in ${init2_files}
do
    echo ${file}
    sshpass -p ${init2_pw} scp -o StrictHostKeyChecking=no ${file} ${init2_id}@${init2_ip}:${init2_fio_conf_dir}
done
echo "prepare init 2 finish"


if [ ${seq_io_time} -gt 1 ]
then
echo "sequential write start"
sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "cd ${init1_fio_conf_dir}; echo ${init1_pw} | sudo -S nohup fio ./sw_tcp_init1.conf > /dev/null 2>&1 &"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "cd ${init2_fio_conf_dir}; echo ${init2_pw} | sudo -S nohup fio ./sw_tcp_init2.conf > /dev/null 2>&1"
echo "sequential write sleep"
sleep 1
fi

if [ ${rand_io_time} -gt 1 ]
then
echo "random write start"
sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "cd ${init1_fio_conf_dir}; echo ${init1_pw} | sudo -S nohup fio ./rw_tcp_init1.conf > /dev/null 2>&1 &"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "cd ${init2_fio_conf_dir}; echo ${init2_pw} | sudo -S nohup fio ./rw_tcp_init2.conf > /dev/null 2>&1"
fi

mkdir -p longterm_fio_result_init_1
mkdir -p longterm_fio_result_init_2
sshpass -p ${init1_pw} scp -o StrictHostKeyChecking=no ${init1_id}@${init1_ip}:${init1_fio_conf_dir}/*.log ./longterm_fio_result_init_1/
sshpass -p ${init2_pw} scp -o StrictHostKeyChecking=no ${init2_id}@${init2_ip}:${init2_fio_conf_dir}/*.log ./longterm_fio_result_init_2/
sshpass -p ${init1_pw} ssh ${init1_id}@${init1_ip} "echo ${init1_pw} | sudo -S rm ${init1_fio_conf_dir}/*.log"
sshpass -p ${init2_pw} ssh ${init2_id}@${init2_ip} "echo ${init2_pw} | sudo -S rm ${init2_fio_conf_dir}/*.log"
sudo ./calculate_file.py


volume_create_and_mount
}

echo "test start"
sudo ./1_start_pos.sh
sudo ./2_bring_up.sh -a ${target_ip_1} -b ${target_ip_2} -s ${volume_cnt} -v ${volume_cnt} -S ${volume_byte_size} -n ${target_nic_1} -m ${target_nic_2}

write_fill_pos_before_vol_delete

longterm_test_after_vol_delete

echo "random write sleep"
