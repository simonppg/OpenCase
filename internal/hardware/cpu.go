package hardware

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// CollectCPU reads CPU information from Linux.
func CollectCPU() (CPUInfo, error) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return CPUInfo{}, err
	}
	defer file.Close()

	info := CPUInfo{
		Architecture: runtime.GOARCH,
		Threads:      runtime.NumCPU(),
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {

		case "model name":
			if info.Model == "" {
				info.Model = value
			}

		case "cpu cores":
			if info.Cores == 0 {
				info.Cores, _ = strconv.Atoi(value)
			}

		case "cpu MHz":
			if info.FrequencyMHz == 0 {
				frequency, _ := strconv.ParseFloat(value, 64)
				info.FrequencyMHz = int(frequency)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return CPUInfo{}, err
	}

	return info, nil
}
