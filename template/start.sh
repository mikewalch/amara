#!/bin/bash

acc_home=./accumulo-{{.AccumuloVersion}}
acc_conf=$acc_home/conf

cp $acc_conf/examples/2GB/standalone/* $acc_conf/
cp accumulo-site.xml $acc_conf/
cp accumulo-env.sh $acc_conf/

$acc_home/bin/accumulo $1
