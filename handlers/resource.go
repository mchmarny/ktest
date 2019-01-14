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
		sr.Node.Resources.Memory.Value = float64(vm.Total)
		sr.Node.Resources.Memory.Context = "Source: OS process status"
	}

	count, err := cpu.Counts(true)
	if err == nil {
		sr.Node.Resources.CPU.Value = float64(count)
		sr.Node.Resources.CPU.Context = "Source: OS process status"
	}

	// pod
	val, ctx := getResourceLimit(limitMemResourceFile)
	sr.Pod.Limits.Memory.Value = val
	sr.Pod.Limits.Memory.Context = ctx

	quotaVal, quotaCtx := getResourceLimit(limitCPUQuotaResourceFile)
	periodVal, _ := getResourceLimit(limitCPUPeriodResourceFile)

	if quotaVal > 0 && periodVal > 0 {
		sr.Pod.Limits.CPU = &types.SimpleIntMeasurement{
			Value: quotaVal / periodVal,
			Context: quotaCtx,
		}
	}else {
		sr.Pod.Limits.CPU.Context = quotaCtx
	}


	writeJSON(w, sr)

}




func getResourceLimit(path string) (val float64, info string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("File[%s]: not found: %v", path, err)
		return 0, fmt.Sprintf("File not found: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file[%s]: %v", path, err)
		return 0, fmt.Sprintf("Unable to open file: %s", path)
	}

	bc, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file[%s]: %v", path, err)
		return 0, fmt.Sprintf("Unable to read file: %s", path)
	}

	cs :=  strings.Trim(string(bc), "\n")

	ic, err := strconv.ParseFloat(cs, 64)
	if err != nil {
		log.Printf("Error parsing content[%s]: %v", path, err)
		return 0, fmt.Sprintf("Non-numeric value: %s", cs)
	}

	log.Printf("File[%s]: %v = %f", path, cs, ic)

	return ic, fmt.Sprintf("Source: %s", path)

}