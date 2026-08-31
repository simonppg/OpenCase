package hardware

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

func CollectSystem() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	cpu, err := CollectCPU()
	if err != nil {
		return SystemInfo{}, err
	}

	memory, err := CollectMemory()
	if err != nil {
		return SystemInfo{}, err
	}

	storage, err := CollectStorage()
	if err != nil {
		return SystemInfo{}, err
	}

	network, err := CollectNetwork()
	if err != nil {
		return SystemInfo{}, err
	}

	pci, err := CollectPCI()
	if err != nil {
		return SystemInfo{}, err
	}

	return SystemInfo{
		Hostname:      hostname,
		OS:            detectOS(),
		Kernel:        detectKernel(),
		Architecture:  runtime.GOARCH,
		UptimeSeconds: collectUptime(),

		CPU:     cpu,
		Memory:  memory,
		Storage: storage,
		Network: network,
		PCI:     pci,
	}, nil
}

func detectOS() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Linux"
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(
				strings.TrimPrefix(line, "PRETTY_NAME="),
				`"`,
			)
		}
	}

	return "Linux"
}

func detectKernel() string {
	data, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "Unknown"
	}

	return strings.TrimSpace(string(data))
}

func collectUptime() uint64 {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0
	}

	fields := strings.Fields(string(data))

	if len(fields) == 0 {
		return 0
	}

	value, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0
	}

	return uint64(value)
}
