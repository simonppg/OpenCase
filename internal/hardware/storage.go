package hardware

import (
	"os"
	"path/filepath"
	"strings"
)

// CollectStorage discovers block devices through sysfs.
func CollectStorage() ([]StorageInfo, error) {
	entries, err := os.ReadDir("/sys/block")
	if err != nil {
		return nil, err
	}

	var devices []StorageInfo

	for _, entry := range entries {
		name := entry.Name()

		// Ignore loop devices.
		if strings.HasPrefix(name, "loop") {
			continue
		}

		sizePath := filepath.Join(
			"/sys/block",
			name,
			"size",
		)

		data, err := os.ReadFile(sizePath)
		if err != nil {
			continue
		}

		sectors := strings.TrimSpace(string(data))

		devices = append(devices, StorageInfo{
			Name:      name,
			Device:    "/dev/" + name,
			Type:      "Block Device",
			SizeBytes: parseSectors(sectors),
		})
	}

	return devices, nil
}

func parseSectors(value string) uint64 {
	var sectors uint64

	for _, char := range value {
		if char < '0' || char > '9' {
			return 0
		}

		sectors = sectors*10 + uint64(char-'0')
	}

	// Linux sysfs reports block-device size in 512-byte sectors.
	return sectors * 512
}
