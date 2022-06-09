package vm

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	firecracker "agent/lib/firecracker"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func copy(src string, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, 0644)
	return err
}

// Create a VMM with a given set of options and start the VM
func createAndStartVM(ctx context.Context) (*MicroVM, error) {
	vmmID := xid.New().String()

	copy("./rootfs.ext4", "/tmp/rootfs-"+vmmID+".ext4")

	fcCfg, err := GetFirecrackerConfig(vmmID)
	if err != nil {
		log.Errorf("Error: %s", err)
		return nil, err
	}

	logger := log.New()

	if false { // TODO
		log.SetLevel(log.DebugLevel)
		logger.SetLevel(log.DebugLevel)
	}

	machineOpts := []firecracker.Opt{
		firecracker.WithLogger(log.NewEntry(logger)),
	}

	firecrackerBinary, err := exec.LookPath("firecracker")
	if err != nil {
		logger.Log(log.ErrorLevel, err)
		return nil, err
	}

	finfo, err := os.Stat(firecrackerBinary)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("binary %q does not exist: %v", firecrackerBinary, err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to stat binary, %q: %v", firecrackerBinary, err)
	}

	if finfo.IsDir() {
		return nil, fmt.Errorf("binary, %q, is a directory", firecrackerBinary)
	} else if finfo.Mode()&0111 == 0 {
		return nil, fmt.Errorf("binary, %q, is not executable. Check permissions of binary", firecrackerBinary)
	}

	// if the jailer is used, the final command will be built in NewMachine()
	if fcCfg.JailerCfg == nil {
		cmd := firecracker.VMCommandBuilder{}.
			WithBin(firecrackerBinary).
			WithSocketPath(fcCfg.SocketPath).
			// WithStdin(os.Stdin).
			// WithStdout(os.Stdout).
			WithStderr(os.Stderr).
			Build(ctx)

		machineOpts = append(machineOpts, firecracker.WithProcessRunner(cmd))
	}

	vmmCtx, vmmCancel := context.WithCancel(ctx)

	m, err := firecracker.NewMachine(vmmCtx, fcCfg, machineOpts...)
	if err != nil {
		vmmCancel()
		logger.Log(log.ErrorLevel, err)
		return nil, fmt.Errorf("failed creating machine: %s", err)
	}

	if err := m.Start(vmmCtx); err != nil {
		vmmCancel()
		logger.Log(log.ErrorLevel, err)
		return nil, fmt.Errorf("failed to start machine: %v", err)
	}

	log.WithField("ip", m.Cfg.NetworkInterfaces[0].StaticConfiguration.IPConfiguration.IPAddr.IP).Info("machine started")

	return &MicroVM{
		Ctx:     vmmCtx,
		Cancel:  vmmCancel,
		ID:      vmmID,
		Machine: m,
		IP:      m.Cfg.NetworkInterfaces[0].StaticConfiguration.IPConfiguration.IPAddr.IP,
	}, nil
}
