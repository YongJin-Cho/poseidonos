#!/bin/bash

ROOT_DIR=$(readlink -f $(dirname $0))/../../
sudo mkdir -p /etc/ibofos/conf
sudo cp ./ibofos.conf /etc/ibofos/conf/ibofos.conf
sudo $ROOT_DIR/script/setup_env.sh
sudo $ROOT_DIR/bin/ibofos > /dev/null &
