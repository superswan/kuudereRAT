#!/bin/bash
cd $3
host_ip=$1 
host_port=$2
echo "\"${host_ip}:${host_port}\""
LHOST="${host_ip}:${host_port}" make
