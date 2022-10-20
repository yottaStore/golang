#!/bin/bash

cd /tmp/
dd if=/dev/zero of=loopbackfile.img bs=16M count=10
losetup -fP loopbackfile.img
mkfs.xfs loopbackfile.img
mount -o loop /dev/loop21 /loopfs/
# Get the loopback device name

chown mamluk /loopfs/