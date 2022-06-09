package web

import (
	"net/http"
	"os/exec"
	"strings"

	v "github.com/klauspost/cpuid/v2"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  bool
	Message string
}

func GetHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetVMInfo(c echo.Context) error {

	type CPUInformation struct {
		Name           string
		Features       string
		Frequency      int64
		Cores          int
		ThreadsPerCore int
		LogicalCores   int
	}
	type VMInfo struct {
		CPU CPUInformation
	}

	return c.JSON(http.StatusOK, &VMInfo{
		CPU: CPUInformation{
			Name:           v.CPU.BrandName,
			Features:       strings.Join(v.CPU.FeatureSet(), " "),
			Frequency:      v.CPU.Hz / 1000000,
			Cores:          v.CPU.PhysicalCores,
			LogicalCores:   v.CPU.LogicalCores,
			ThreadsPerCore: v.CPU.ThreadsPerCore,
		},
	})
}

func RunFunction(c echo.Context) error {
	packVersion := exec.Command("docker", "version")

	res, _ := packVersion.Output()

	return c.JSON(http.StatusOK, Response{
		Status:  true,
		Message: string(res),
	})
}
