package vm

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"agent/lib/firecracker"
	"agent/lib/firecracker/client/models"
)

func GetFirecrackerConfig(vmmID string) (firecracker.Config, error) {
	socket := GetSocketPath(vmmID)
	return firecracker.Config{
		SocketPath:      socket,
		KernelImagePath: "./vmlinux",
		LogPath:         fmt.Sprintf("%s.log", socket),
		Drives: []models.Drive{{
			DriveID:      firecracker.String("0"),
			PathOnHost:   firecracker.String("/tmp/rootfs-" + vmmID + ".ext4"),
			IsRootDevice: firecracker.Bool(true),
			IsReadOnly:   firecracker.Bool(false),
			RateLimiter: firecracker.NewRateLimiter(
				// bytes/s
				models.TokenBucket{
					OneTimeBurst: firecracker.Int64(1024 * 1024), // 1 MiB/s
					RefillTime:   firecracker.Int64(500),         // 0.5s
					Size:         firecracker.Int64(1024 * 1024),
				},
				// ops/s
				models.TokenBucket{
					OneTimeBurst: firecracker.Int64(100),  // 100 iops
					RefillTime:   firecracker.Int64(1000), // 1s
					Size:         firecracker.Int64(100),
				}),
		}},
		NetworkInterfaces: []firecracker.NetworkInterface{{
			// Use CNI to get dynamic IP
			CNIConfiguration: &firecracker.CNIConfiguration{
				NetworkName: "fcnet",
				IfName:      "veth0",
			},
		}},
		MachineCfg: models.MachineConfiguration{
			VcpuCount:  firecracker.Int64(1),
			MemSizeMib: firecracker.Int64(128),
		},
	}, nil
}

func GetSocketPath(vmmID string) string {
	filename := strings.Join([]string{
		".firecracker.sock",
		strconv.Itoa(os.Getpid()),
		vmmID,
	},
		"-",
	)
	dir := os.TempDir()

	return filepath.Join(dir, filename)
}
