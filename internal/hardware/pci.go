package hardware

import (
	"os"
	"path/filepath"
	"strings"
)

func CollectPCI() ([]PCIInfo, error) {
	const pciPath = "/sys/bus/pci/devices"

	entries, err := os.ReadDir(pciPath)
	if err != nil {
		return nil, err
	}

	devices := make([]PCIInfo, 0, len(entries))

	for _, entry := range entries {
		address := entry.Name()
		devicePath := filepath.Join(pciPath, address)

		device := PCIInfo{
			Address:  address,
			VendorID: readSysfs(filepath.Join(devicePath, "vendor")),
			DeviceID: readSysfs(filepath.Join(devicePath, "device")),
			Class:    readSysfs(filepath.Join(devicePath, "class")),
		}

		driverPath := filepath.Join(devicePath, "driver")

		if driver, err := os.Readlink(driverPath); err == nil {
			device.Driver = filepath.Base(driver)
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func readSysfs(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}
