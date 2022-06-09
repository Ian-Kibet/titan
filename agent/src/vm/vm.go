package vm

import (
	"context"
	"net"
	"os"
	"time"

	"agent/lib/firecracker"

	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

type MicroVM struct {
	Ctx     context.Context
	Cancel  context.CancelFunc
	ID      string
	Machine *firecracker.Machine
	IP      net.IP
}

func waitForVMToBoot(ctx context.Context, ip net.IP) error {
	// Query the agent until it provides a valid response
	req.SetTimeout(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			// Timeout
			return ctx.Err()
		default:
			res, err := req.Get("http://" + ip.String() + ":8080/health")
			if err != nil {
				log.WithError(err).Info("VM not ready yet")
				time.Sleep(time.Second)
				continue
			}

			if res.Response().StatusCode != 200 {
				time.Sleep(time.Second)
				log.Info("VM not ready yet")
			} else {
				log.WithField("ip", ip).Info("VM agent ready")
				return nil
			}
			time.Sleep(time.Second)
		}

	}
}

func (vm MicroVM) ShutDown() {
	log.WithField("ip", vm.IP).Info("stopping")
	vm.Machine.StopVMM()
	err := os.Remove(vm.Machine.Cfg.SocketPath)
	if err != nil {
		log.WithError(err).Error("Failed to delete firecracker socket")
	}
	err = os.Remove("/tmp/rootfs-" + vm.ID + ".ext4")
	if err != nil {
		log.WithError(err).Error("Failed to delete firecracker rootfs")
	}
}

func FillVMPool(ctx context.Context, WarmVMs chan<- MicroVM) {
	for {
		select {
		case <-ctx.Done():
			// Program is stopping, WarmVMs will be cleaned up, bye
			return
		default:
			vm, err := createAndStartVM(ctx)
			if err != nil {
				log.Error("failed to create VMM")
				time.Sleep(time.Second)
				continue
			}

			log.WithField("ip", vm.IP).Info("New VM created and started")

			// Don't wait forever, if the VM is not available after 10s, move on
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			err = waitForVMToBoot(ctx, vm.IP)
			if err != nil {
				log.WithError(err).Info("VM not ready yet")
				vm.Cancel()
				continue
			}

			// Add the new microVM to the pool.
			// If the pool is full, this line will block until a slot is available.
			WarmVMs <- *vm

			log.Info("VMs Available", len(WarmVMs))
		}
	}
}
