#!/bin/sh

set -eu

arch=`uname -m`
dest_kernel="vmlinux"
image_bucket_url="https://s3.amazonaws.com/spec.ccfc.min/img/quickstart_guide/$arch"

if [ ${arch} = "x86_64" ]; then
    kernel="${image_bucket_url}/kernels/vmlinux.bin"
elif [ ${arch} = "aarch64" ]; then
    kernel="${image_bucket_url}/kernels/vmlinux.bin"
else
    echo "Cannot run firecracker on $arch architecture!"
    exit 1
fi

if [ ! -f $dest_kernel ]; then
    echo "Kernel not found, downloading $kernel..."
    curl --progress-bar -fsSL -o $dest_kernel $kernel --progress-bar
    echo "Saved kernel file to $dest_kernel."
else
    echo "Kernel already exists in your directory"
fi