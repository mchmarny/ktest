package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"
)

func HostHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Host...")

	var request []string

	info, infoErr := host.Info()

	if infoErr == nil {
		request = append(request, fmt.Sprintf("Node Host ID: %v", info.HostID))
		request = append(request, fmt.Sprintf("   Boot time: %v", time.Unix(int64(info.BootTime), 0)))
		request = append(request, fmt.Sprintf("   Hostname: %v", info.Hostname))
		request = append(request, fmt.Sprintf("   OS: %v", info.OS))
	}

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
