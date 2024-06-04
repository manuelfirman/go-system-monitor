package application

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/manuelfirman/go-system-monitor/internal/hardware"
	"github.com/manuelfirman/go-system-monitor/internal/server"
)

type ApplicationDefault struct {
	server *server.Default
}

// NewApplicationDefault is the function that creates a new ApplicationDefault.
func NewApplicationDefault(s *server.Default) *ApplicationDefault {
	return &ApplicationDefault{
		server: s,
	}
}

// SetUp is the method that sets up the application.
func (s *ApplicationDefault) Listen() error {
	return http.ListenAndServe(":8080", &s.server.Mux)
}

// Run is the method that runs the application.
func (s *ApplicationDefault) Run() {
	func(srv *server.Default) {
		for {
			systemData, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
				continue
			}
			diskData, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println(err)
				continue
			}
			cpuData, err := hardware.GetCPUSection()
			if err != nil {
				fmt.Println(err)
				continue
			}
			timeStamp := time.Now().Format("2006-01-02 15:04:05")

			fmt.Println("System data:", systemData)
			fmt.Println("Disk data:", diskData)
			fmt.Println("CPU data:", cpuData)
			fmt.Println("Timestamp:", timeStamp)
			time.Sleep(3 * time.Second)
		}
	}(s.server)
}

// serveSystemHtml serves the system data in HTML format.
func serveSystemHtml(s *hardware.SystemSection) (html string) {
	html = "<div class='system-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html = html + "<tr><td>Operating System:</td> <td><i class='fa fa-brands fa-linux'></i> " + s.RuntimeOS + "</td></tr>"
	html = html + "<tr><td>Platform:</td><td> <i class='fa fa-brands fa-fedora'></i> " + s.Platform + "</td></tr>"
	html = html + "<tr><td>Hostname:</td><td>" + s.Hostname + "</td></tr>"
	html = html + "<tr><td>Number of processes running:</td><td>" + strconv.FormatUint(s.Procs, 10) + "</td></tr>"
	html = html + "<tr><td>Total memory:</td><td>" + strconv.FormatUint(s.TotalVM, 10) + " MB</td></tr>"
	html = html + "<tr><td>Free memory:</td><td>" + strconv.FormatUint(s.FreeVM, 10) + " MB</td></tr>"
	html = html + "<tr><td>Percentage used memory:</td><td>" + strconv.FormatFloat(s.PercentageMemoryUsed, 'f', 2, 64) + "%</td></tr></tbody></table>"

	html = html + "</div>"

	return
}

// serveDiskHtml serves the disk data in HTML format.
func serveCPUHtml(c *hardware.CPUSection) (html string) {
	html = "<div class='cpu-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html = html + "<tr><td>Model Name:</td><td>" + c.ModelName + "</td></tr>"
	html = html + "<tr><td>Family:</td><td>" + c.Family + "</td></tr>"
	html = html + "<tr><td>Speed:</td><td>" + strconv.FormatFloat(c.Speed, 'f', 2, 64) + " MHz</td></tr>"

	html = html + "<tr><td>Cores: </td><td><div class='row mb-4'><div class='col-md-6'><table class='table table-sm'><tbody>"

	// iterate over the first half of the CPU percentages
	for idx, cpupercent := range c.FirstCPU {
		html = html + "<tr><td>CPU [" + strconv.Itoa(idx) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%</td></tr>"
	}
	html = html + "</tbody></table></div><div class='col-md-6'><table class='table table-sm'><tbody>"

	// iterate over the second half of the CPU percentages
	for idx, cpupercent := range c.SecondCPU {
		html = html + "<tr><td>CPU [" + strconv.Itoa(idx+8) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%</td></tr>"
	}

	html = html + "</tbody></table></div></div></td></tr></tbody></table></div>"
	return
}

// serveDiskHtml serves the disk data in HTML format.
func serveDiskHtml(d *hardware.DiskSection) (html string) {
	html = "<div class='disk-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html = html + "<tr><td>Total disk space:</td><td>" + strconv.FormatUint(d.TotalSpace, 10) + " GB</td></tr>"
	html = html + "<tr><td>Used disk space:</td><td>" + strconv.FormatUint(d.UsedSpace, 10) + " GB</td></tr>"
	html = html + "<tr><td>Free disk space:</td><td>" + strconv.FormatUint(d.FreeSpace, 10) + " GB</td></tr>"
	html = html + "<tr><td>Percentage disk space usage:</td><td>" + strconv.FormatFloat(d.PercentageUsed, 'f', 2, 64) + "%</td></tr>"

	return
}
