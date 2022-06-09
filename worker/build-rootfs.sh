#!/bin/bash

set -xe

echo ""
echo "Starting build"

[ "$UID" -eq 0 ] || { echo "Root permission is required"; exit 1;}

fs_dest=../agent/rootfs.ext4

dd if=/dev/zero of=$fs_dest bs=1M count=1000
mkfs.ext4 $fs_dest
mkdir -p /tmp/my-rootfs
mount $fs_dest /tmp/my-rootfs

docker run -i --rm \
    -v /tmp/my-rootfs:/my-rootfs \
    -v "$(pwd)/worker:/usr/local/bin/worker" \
    -v "$(pwd)/openrc-service.sh:/etc/init.d/worker" \
    alpine sh <setup-alpine.sh

umount /tmp/my-rootfs
e2fsck -f $fs_dest
resize2fs -M $fs_dest

echo ""
echo "File system successfully built and saved in $fs_dest"