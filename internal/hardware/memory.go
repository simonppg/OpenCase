package hardware

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// CollectMemory reads memory information from Linux.
func CollectMemory() (MemoryInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return MemoryInfo{}, err
	}
	defer file.Close()

	var total uint64
	var available uint64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) < 2 {
			continue
		}

		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}

		// /proc/meminfo reports values in KiB.
		value *= 1024

		switch fields[0] {

		case "MemTotal:":
			total = value

		case "MemAvailable:":
			available = value
		}
	}

	if err := scanner.Err(); err != nil {
		return MemoryInfo{}, err
	}

	return MemoryInfo{
		TotalBytes:     total,
		AvailableBytes: available,
		UsedBytes:      total - available,
	}, nil
}
