package main

import (
	"context"
	"fmt"
	"os"
	firecracker "titan/firecracker"
	models "titan/lib/firecracker/client/models"

	log "github.com/sirupsen/logrus"
	// firecracker "titan/firecracker"
)

const (
	// How Firecracker is launched
	firecrackerPath = "./firecracker"
	kernelPath      = "vmlinux.bin"
	rootfsPath      = "rootfs.ext4"

	// Firecracker settings
	noCpus                 = 2
	memorySize             = 4096
	kernelArgs             = "console=ttyS0 reboot=k panic=1 pci=off quiet"
	firecrackerInitTimeout = 3
)

func launchVM(socketPath string) {
	// Remove the socket path if it exists
	if _, err := os.Stat(socketPath); err == nil {
		os.Remove(socketPath)
	}

	// Create a config structure that specifies how we launch
	// the microVM.
	cfg := firecracker.Config{
		SocketPath:      socketPath,
		KernelImagePath: kernelPath,
		KernelArgs:      kernelArgs,
		Drives:          firecracker.NewDrivesBuilder(rootfsPath).Build(),
		MachineCfg: models.MachineConfiguration{
			VcpuCount:  firecracker.Int64(noCpus),
			MemSizeMib: firecracker.Int64(memorySize),
		},
	}

	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Build the command
	cmd := firecracker.VMCommandBuilder{}.
		WithSocketPath(socketPath).
		WithBin(firecrackerPath).
		WithStdin(os.Stdin).
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		Build(ctx)

	// Create a logger to have a nice output
	logger := log.New()

	// Create the machine instance
	machine, err := firecracker.NewMachine(
		ctx,
		cfg,
		firecracker.WithProcessRunner(cmd),
		firecracker.WithLogger(log.NewEntry(logger)))

	if err != nil {
		panic(fmt.Errorf("failed to create new machine: %v", err))
	}

	// Start the microVM
	if err := machine.Start(ctx); err != nil {
		panic(fmt.Errorf("Failed to start machine: %v", err))
	}
	defer machine.StopVMM()

	// wait for the VMM to exit
	if err := machine.Wait(ctx); err != nil {
		panic(fmt.Errorf("Wait returned an error %s", err))
	}
	os.Remove(socketPath)
}

func main() {

	fmt.Println("hfgfhgg")
	// socketPath := flag.String("socket", "", "UDS socket path for Firecracker to use.")
	// flag.Parse()

	// if *socketPath == "" {
	// 	panic(fmt.Errorf("UDS socket path needed."))
	// }

	// launchVM(*socketPath)
}
