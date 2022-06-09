### Primelab CPU Agent

TODO: Describe this software

## Getting Started

What you need to get started

[] [golang](https://go.dev) v1.16.6 - v1.18.x 
[] vmlinux (Use the init-vmlinux.sh script to fetch)
   ```sh
   sudo ./scripts/init-vmlinux.sh
   ```
[] firecracker (Use the init-firecracker.sh script to fetch)
   ```sh
   sudo ./scripts/init-firecracker.sh
   ```
[] rootfs.ext4 
   ```sh
   sudo ../worker/build-rootfs.sh
   ```


## Build

```sh
go build
```
This will fetch all required packages and build the binary `agent`

## Run it 

Prebuilt binary

```sh
./agent
```

Code

```sh
go run .
```


Need more help, slack me -> Ian Kibet



