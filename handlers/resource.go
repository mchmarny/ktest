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
	"github.com/mchmarny/tellmeall/utils"

	"golang.org/x/sys/unix"

	"github.com/jaypipes/ghw"
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
		Node: getNodeInfo(),
		Pod: getPodInfo(),
	}

	writeJSON(w, sr)

}

func getPodInfo() *types.SimplePodInfo {

	pod := types.NewPodInfo()

	// host
	info, err := host.Info()
	if err == nil {
		pod.Hostname = info.Hostname
	}

	// pod memory
	val, wr, ctx := getCGroupsFile(limitMemResourceFile)
	pod.Limits.RAM.Value = val
	pod.Limits.RAM.Context = fmt.Sprintf("%s, Writable: %v, Size: %s", ctx, wr,
		utils.ByteSize(uint64(val)))

	// pod cpu (calculated: quota / period)
	quotaVal, wr, quotaCtx := getCGroupsFile(limitCPUQuotaResourceFile)
	periodVal, _, _ := getCGroupsFile(limitCPUPeriodResourceFile)

	// always context
	pod.Limits.CPU.Context = fmt.Sprintf("%s, Writable: %v", quotaCtx, wr)

	// value only if there is data
	if quotaVal > 0 && periodVal > 0 {
		pod.Limits.CPU.Value = quotaVal / periodVal
	}

	return pod
}


func getNodeInfo() *types.SimpleNodeInfo {

	node := types.NewNodeInfo()

	// host
	info, err := host.Info()
	if err == nil {
		node.ID = info.HostID
		node.BootTime = time.Unix(int64(info.BootTime), 0)
		node.OS = info.OS
	}

	// vm
	vm, err := mem.VirtualMemory()
	if err == nil {
		node.Resources.RAM.Value = float64(vm.Total)
		node.Resources.RAM.Context = fmt.Sprintf(
			"Source: OS process status, Size: %s",
			utils.ByteSize(vm.Total))
	}

	// cpu
	count, err := cpu.Counts(true)
	if err == nil {
		node.Resources.CPU.Value = float64(count)
		node.Resources.CPU.Context = "Source: OS process status"
	}

	// gpu
	gpuCount, deviceInfo := getGPUInfo()
	if gpuCount > 0 {
		node.Resources.GPU = &types.SimpleMeasurement{
			Value: float64(gpuCount),
			Context: deviceInfo,
		}
	}

	return node

}



func getCGroupsFile(path string) (val float64, wr bool, info string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("File[%s]: not found: %v", path, err)
		return 0, false, fmt.Sprintf("File not found: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file[%s]: %v", path, err)
		return 0, false, fmt.Sprintf("Unable to open file: %s", path)
	}

	bc, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file[%s]: %v", path, err)
		return 0, false, fmt.Sprintf("Unable to read file: %s", path)
	}

	cs :=  strings.Trim(string(bc), "\n")

	ic, err := strconv.ParseFloat(cs, 64)
	if err != nil {
		log.Printf("Error parsing content[%s]: %v", path, err)
		return 0, false, fmt.Sprintf("Non-numeric value: %s", cs)
	}

	log.Printf("File[%s]: %v = %f", path, cs, ic)

	return ic, unix.Access(path, unix.W_OK) == nil, fmt.Sprintf("Source: %s", path)

}




func getGPUInfo() (val int, info string) {

	gpu, err := ghw.GPU()
	if err != nil || gpu == nil ||  gpu.GraphicsCards == nil {
		log.Printf("Error getting GPU info: %v", err)
		return 0, ""
	}

	num := len(gpu.GraphicsCards)
	infos := make([]string, 0)

	for _, card := range gpu.GraphicsCards {
		log.Printf(" %v\n", card)
		infos = append(infos, fmt.Sprintf("gpu[%d]: %s - %s",
			card.Index, card.DeviceInfo.Vendor.Name, card.DeviceInfo.Product.Name))
	}

	return num, strings.Join(infos, "; ")

}