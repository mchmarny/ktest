package handlers

import (
	"log"
	"net/http"
	"time"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"fmt"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"

	"github.com/shirou/gopsutil/host"

	"github.com/mchmarny/tellmeall/types"
)

const (
	limitMemResourceFile       = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	limitCPUPeriodResourceFile = "/sys/fs/cgroup/cpu/cpu.cfs_period_us"
	limitCPUQuotaResourceFile  = "/sys/fs/cgroup/cpu/cpu.cfs_quota_us"
)

func resourceHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Resource...")

	w.Header().Set("Content-Type", "application/json")

	sr := &types.SimpleResource{
		Meta: getMeta(r),
		Node: &types.SimpleNodeInfo{
			Resources: &types.SimpleResourceInfo{
				Memory: &types.SimpleIntMeasurement{},
				CPU: &types.SimpleIntMeasurement{},
			},
		},
		Pod: &types.SimplePodInfo{
			Limits:   &types.SimpleResourceInfo{
				Memory: &types.SimpleIntMeasurement{},
				CPU: &types.SimpleIntMeasurement{},
			},
			Current: &types.SimpleResourceInfo{
				CPU: &types.SimpleIntMeasurement{},
			},
		},
	}

	info, err := host.Info()
	if err == nil {
		sr.Node.ID = info.HostID
		sr.Node.BootTime = time.Unix(int64(info.BootTime), 0)
		sr.Node.OS = info.OS
		sr.Pod.Hostname = info.Hostname
	}


	vm, err := mem.VirtualMemory()
	if err == nil {
		sr.Node.Resources.Memory.Value = int64(vm.Total)
		sr.Node.Resources.Memory.Context = "ps"
	}

	count, err := cpu.Counts(true)
	if err == nil {
		sr.Node.Resources.CPU.Value = int64(count)
		sr.Node.Resources.CPU.Context = "ps"
	}

	// pod
	sr.Pod.Limits.Memory = getResourceLimit(limitMemResourceFile)
	sr.Pod.Limits.CPU = getResourceLimit(limitCPUQuotaResourceFile)
	sr.Pod.Current.CPU = getResourceLimit(limitCPUPeriodResourceFile)


	writeJSON(w, sr)

}




func getResourceLimit(path string) *types.SimpleIntMeasurement {

	sm := &types.SimpleIntMeasurement{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("File[%s]: not found: %v", path, err)
		sm.Context = fmt.Sprintf("File not found: %s", path)
		return sm
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file[%s]: %v", path, err)
		sm.Context = fmt.Sprintf("Unable to open file: %s", path)
		return sm
	}

	bc, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file[%s]: %v", path, err)
		sm.Context = fmt.Sprintf("Unable to read file: %s", path)
		return sm
	}

	cs :=  strings.Trim(string(bc), "\n")

	ic, err := strconv.ParseInt(cs, 10, 64)
	if err != nil {
		log.Printf("Error parsing content[%s]: %v", path, err)
		sm.Context = fmt.Sprintf("Non-numeric value: %s", cs)
		return sm
	}

	log.Printf("File[%s]: %v = %d", path, cs, ic)

	sm.Value = ic
	sm.Context = fmt.Sprintf("Source: %s", path)

	return sm

}