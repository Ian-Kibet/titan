package main

import (
	http "agent/src/http"
	"agent/src/vm"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

// const (
// 	// executableMask is the mask needed to check whether or not a file's
// 	// permissions are executable.
// 	executableMask = 0111

// 	// How Firecracker is launched
// 	firecrackerPath = "./firecracker"
// 	kernelPath      = "vmlinux.bin"
// 	rootfsPath      = "rootfs.ext4"

// 	// Firecracker settings
// 	noCpus                 = 2
// 	memorySize             = 128
// 	kernelArgs             = "console=ttyS0 reboot=k panic=1 pci=off quiet"
// 	firecrackerInitTimeout = 3
// )

func installSignalHandlers() {
	go func() {
		// Clear some default handlers installed by the firecracker SDK:
		signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		for {
			switch s := <-c; {
			case s == syscall.SIGTERM || s == os.Interrupt:
				log.Printf("Caught signal: %s, requesting clean shutdown", s.String())
				deleteVMMSockets()
				os.Exit(0)
			case s == syscall.SIGQUIT:
				log.Printf("Caught signal: %s, forcing shutdown", s.String())
				deleteVMMSockets()
				os.Exit(0)
			}
		}
	}()
}

func deleteVMMSockets() {
	log.Debug("cc")
	dir, err := ioutil.ReadDir(os.TempDir())
	if err != nil {
		log.WithError(err).Error("Failed to read directory")
	}
	for _, d := range dir {
		log.WithField("d", d.Name()).Debug("considering")
		if strings.Contains(d.Name(), fmt.Sprintf(".firecracker.sock-%d-", os.Getpid())) {
			log.WithField("d", d.Name()).Debug("should delete")
			os.Remove(path.Join([]string{"tmp", d.Name()}...))
		}
	}
}

func main() {
	defer deleteVMMSockets()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	WarmVMs := make(chan vm.MicroVM, 1)

	go vm.FillVMPool(ctx, WarmVMs)
	installSignalHandlers()
	log.SetReportCaller(true)

	vm := <-WarmVMs

	fmt.Println("VM IP ", vm.IP)

	http.StartHttpServer("8081")

	//Defer cleanup of VM and VMM
	go func() {
		defer vm.Cancel()
		vm.Machine.Wait(vm.Ctx)
		time.Sleep(time.Minute)
	}()
	defer vm.ShutDown()
}
