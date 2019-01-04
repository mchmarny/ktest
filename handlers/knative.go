package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	runtimeContractURL = `Knative Runtime Contract:
	https://github.com/knative/serving/blob/master/docs/runtime-contract.md`
)

func isDirRW(path string) string {

	if unix.Access(path, unix.O_RDWR) == nil {
		return "R/W (Success)"
	}

	if unix.Access(path, unix.O_RDONLY) == nil {
		return "R/- (Failed: no write)"
	}

	if unix.Access(path, unix.O_WRONLY) == nil {
		return "-/W (Failed: no read)"
	}

	return "-/- (Failed: neither)"
}

func knativeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Knative...")

	var request []string

	request = append(request, runtimeContractURL)
	request = append(request, "\n")

	request = append(request, "=== ENVIRONMENT VARIABLES ===")
	request = append(request, fmt.Sprintf("   PORT:            %v", os.Getenv("PORT")))
	request = append(request, fmt.Sprintf("   K_SERVICE:       %v", os.Getenv("K_SERVICE")))
	request = append(request, fmt.Sprintf("   K_REVISION:      %v", os.Getenv("K_REVISION")))
	request = append(request, fmt.Sprintf("   K_CONFIGURATION: %v", os.Getenv("K_CONFIGURATION")))
	request = append(request, "\n")

	request = append(request, "=== FILESYSTEM (Required R/W) ===")
	request = append(request, fmt.Sprintf("/tmp     - %s", isDirRW("/tmp")))
	request = append(request, fmt.Sprintf("/var/log - %s", isDirRW("/var/log")))
	request = append(request, fmt.Sprintf("/dev/log - %s", isDirRW("/dev/log")))
	request = append(request, "\n")

	request = append(request, "=== DNS (Optional) ===")
	request = append(request, fmt.Sprintf("/etc/hosts       - %s", isDirRW("/etc/hosts")))
	request = append(request, fmt.Sprintf("/etc/hostname    - %s", isDirRW("/etc/hostname")))
	request = append(request, fmt.Sprintf("/etc/resolv.conf - %s", isDirRW("/etc/resolv.conf")))
	request = append(request, "\n")

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
