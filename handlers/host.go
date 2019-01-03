package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/docker"
)

func HostHandler(w http.ResponseWriter, r *http.Request) {

	var request []string

	info, infoErr := host.Info()
	docIDs, docErr := docker.GetDockerIDList()

	if infoErr == nil {
		request = append(request, fmt.Sprintf("Host ID: %v", info.HostID))
		request = append(request, fmt.Sprintf("Hostname: %v", info.Hostname))
		request = append(request, fmt.Sprintf("Platform: %v", info.PlatformVersion))
		request = append(request, fmt.Sprintf("OS: %v", info.OS))
		request = append(request, fmt.Sprintf("Boot-time: %v", time.Unix(int64(info.BootTime), 0)))
		request = append(request, fmt.Sprintf("Uptime: %v", time.Duration(info.Uptime)))
	}

	if docErr == nil {
		request = append(request, fmt.Sprintf("Containers: %v", strings.Join(docIDs, ", ")))
	}

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
