#!/bin/bash

acc_home=./accumulo-1.7.1
acc_conf=$acc_home/conf

#$acc_home/bin/bootstrap_config.sh

cp $acc_conf/examples/2GB/standalone/* $acc_conf/
cp accumulo-site.xml $acc_conf/
cp accumulo-env.sh $acc_conf/

$acc_home/bin/accumulo $1
