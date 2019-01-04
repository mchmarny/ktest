package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shirou/gopsutil/mem"
)

func memoryHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Memory...")

	var request []string

	v, _ := mem.VirtualMemory()

	request = append(request, fmt.Sprintf("Total: %d", v.Total))
	request = append(request, fmt.Sprintf("Free:  %d", v.Free))
	request = append(request, fmt.Sprintf("Used:  %.2f percent", v.UsedPercent))

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
