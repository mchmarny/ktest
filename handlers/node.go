package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/host"

	"github.com/mchmarny/tellmeall/types"
)

func nodeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Node...")

	w.Header().Set("Content-Type", "application/json")

	sr := &types.SimpleNode{
		Meta: getMeta(r),
		Info: &types.SimpleNodeInfo{
			Memory: &types.SimpleMemory{},
		},
	}

	vm, err := mem.VirtualMemory()
	if err == nil {
		sr.Info.Memory.Total = vm.Total
		sr.Info.Memory.Free = vm.Free
		sr.Info.Memory.UsedPercent = vm.UsedPercent
	}

	info, err := host.Info()
	if err == nil {
		sr.Info.ID = info.HostID
		sr.Info.BootTime = time.Unix(int64(info.BootTime), 0)
		sr.Info.OS = info.OS
		sr.Info.Hostname = info.Hostname
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sr)

}
