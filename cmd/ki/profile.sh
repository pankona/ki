#!/bin/bash -e

#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 0 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 1 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 2 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 3 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 4 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 5 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 10 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 20 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 30 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 40 ../../../..
#sudo sh -c "echo 1 > /proc/sys/vm/drop_caches" && ./ki -with-profile -c 50 ../../../..

# make page cache
./ki -c 0 ../../../.. &>/dev/null

./ki -with-profile -c 0 ../../../..
./ki -with-profile -c 1 ../../../..
./ki -with-profile -c 2 ../../../..
./ki -with-profile -c 3 ../../../..
./ki -with-profile -c 4 ../../../..
./ki -with-profile -c 5 ../../../..
./ki -with-profile -c 10 ../../../..
./ki -with-profile -c 20 ../../../..
./ki -with-profile -c 30 ../../../..
./ki -with-profile -c 40 ../../../..
./ki -with-profile -c 50 ../../../..
