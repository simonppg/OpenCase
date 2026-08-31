package hardware

type CPUInfo struct {
	Model        string
	Architecture string
	Cores        int
	Threads      int
	FrequencyMHz int
}

type MemoryInfo struct {
	TotalBytes     uint64
	AvailableBytes uint64
	UsedBytes      uint64
}

type StorageInfo struct {
	Name      string
	Device    string
	Type      string
	SizeBytes uint64
}

type NetworkInfo struct {
	Name       string
	State      string
	MACAddress string
	Addresses  []string
}

type PCIInfo struct {
	Address  string
	VendorID string
	DeviceID string
	Class    string
	Driver   string
}

type SystemInfo struct {
	Hostname      string
	OS            string
	Kernel        string
	Architecture  string
	UptimeSeconds uint64

	CPU     CPUInfo
	Memory  MemoryInfo
	Storage []StorageInfo
	Network []NetworkInfo
	PCI     []PCIInfo
}
