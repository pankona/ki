#!/bin/bash -e

#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 0 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 1 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 2 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 3 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 4 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 5 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 10 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 20 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 30 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 40 ../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -c 50 ../..

# make page cache
./ki -c 0 ../.. &>/dev/null

./ki -c 0 ../..
./ki -c 1 ../..
./ki -c 2 ../..
./ki -c 3 ../..
./ki -c 4 ../..
./ki -c 5 ../..
./ki -c 10 ../..
./ki -c 20 ../..
./ki -c 30 ../..
./ki -c 40 ../..
./ki -c 50 ../..
