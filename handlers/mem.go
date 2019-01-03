package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shirou/gopsutil/mem"
)

func MemoryHandler(w http.ResponseWriter, r *http.Request) {

	var request []string

	v, _ := mem.VirtualMemory()

	request = append(request, fmt.Sprintf("Total: %d", v.Total))
	request = append(request, fmt.Sprintf("Free:  %d", v.Free))
	request = append(request, fmt.Sprintf("Used:  %.2f percent", v.UsedPercent))

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
