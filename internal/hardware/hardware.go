package hardware

// SystemSection is the structure for the system section.
type SystemSection struct {
	RuntimeOS     string
	VirtualMemory uint64
	Host          string
}

// CPUSection is the structure for the CPU section.
type CPUSection struct {
	ModelName string
	Cores     int
}

// DiskSection is the structure for the disk section.
type DiskSection struct {
	TotalSpace uint64
	FreeSpace  uint64
	UsedSpace  uint64
}

// HardwareMonitor is the interface that wraps the basic methods of a hardware monitor.
type HardwareMonitor interface {
	// GetSystemSection returns the system section.
	GetSystemSection() (SystemSection, error)
	// GetCPUSection returns the CPU section.
	GetCPUSection() (CPUSection, error)
	// GetDiskSection returns the disk section.
	GetDiskSection() (DiskSection, error)
}
