package hardware

import "errors"

var (
	// ErrNoCPUInfo is the error for no CPU info.
	ErrNoCPUInfo = errors.New("no cpu info")
)

// SystemSection is the structure for the system section.
type SystemSection struct {
	RuntimeOS string
	TotalVM   uint64
	UsedVM    uint64
	Host      string
}

// CPUSection is the structure for the CPU section.
type CPUSection struct {
	ModelName string
	Family    string
	Speed     float64
	FirstCPU  []float64
	SecondCPU []float64
}

// DiskSection is the structure for the disk section.
type DiskSection struct {
	TotalSpace     uint64
	FreeSpace      uint64
	UsedSpace      uint64
	PercentageUsed float64
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
