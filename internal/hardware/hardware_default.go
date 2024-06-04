package hardware

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// megaByteDiv is the division of a megabyte.
const megabyteDiv uint64 = 1024 * 1024

// gigabyteDiv is the division of a gigabyte.
const gigabyteDiv uint64 = megabyteDiv * 1024

// SystemSection is the structure for the system section.
func GetSystemSection() (s SystemSection, err error) {
	// Get the runtime OS.
	s.RuntimeOS = runtime.GOOS

	// Get the virtual memory stats.
	vmStats, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	s.TotalVM = vmStats.Total / megabyteDiv
	s.UsedVM = vmStats.Used / megabyteDiv

	// Get the host info.
	hostInfo, err := host.Info()
	if err != nil {
		return
	}
	s.Host = hostInfo.Hostname

	return
}

// GetCPUSection returns the CPU section.
func GetCPUSection() (c CPUSection, err error) {
	// Get the CPU info.
	cpuStat, err := cpu.Info()
	if err != nil {
		return
	}
	if len(cpuStat) == 0 {
		err = ErrNoCPUInfo
		return
	}
	// Set the CPU info if there is any.
	c.ModelName = cpuStat[0].ModelName
	c.Family = cpuStat[0].Family
	c.Speed = float64(cpuStat[0].Mhz)

	// Get the CPU percentage.
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return
	}
	// Set the CPU percentage.
	c.FirstCPU = percentage[:len(percentage)/2]
	c.SecondCPU = percentage[len(percentage)/2:]

	return
}

// GetDiskSection returns the disk section.
func GetDiskSection() (d DiskSection, err error) {
	// Get the disk stats.
	diskStat, err := disk.Usage("/")
	if err != nil {
		return
	}

	// Set the disk stats.
	d.TotalSpace = diskStat.Total / gigabyteDiv
	d.FreeSpace = diskStat.Free / gigabyteDiv
	d.UsedSpace = diskStat.Used / gigabyteDiv
	d.PercentageUsed = diskStat.UsedPercent

	return
}
