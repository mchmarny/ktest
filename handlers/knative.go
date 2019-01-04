package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/sys/unix"

	"github.com/mchmarny/tellmeall/types"
)

const (
	runtimeContractURL = "https://github.com/knative/serving/blob/master/docs/runtime-contract.md"
)

func isDirRW(path string) *types.FsAccessInfo {

	info := &types.FsAccessInfo{
		Path:     path,
		Expected: "R/W",
	}

	if unix.Access(path, unix.O_RDWR) == nil {
		info.Actual = "R/W"
		info.Comment = "(Success)"
		return info
	}

	if unix.Access(path, unix.O_RDONLY) == nil {
		info.Actual = "R/-"
		info.Comment = "(Failed: no write)"
		return info
	}

	if unix.Access(path, unix.O_WRONLY) == nil {
		info.Actual = "-/W"
		info.Comment = "(Failed: no read)"
		return info
	}

	info.Actual = "-/-"
	info.Comment = "(Failed: neither)"
	return info

}

func knativeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Knative...")
	w.Header().Set("Content-Type", "application/json")

	k := &types.Knative{
		Meta:    getMeta(r),
		InfoURI: runtimeContractURL,
		EnvVars: make(map[string]interface{}),
		Access:  make([]*types.FsAccessGroup, 0),
	}

	// env vars
	k.EnvVars["PORT"] = os.Getenv("PORT")
	k.EnvVars["K_SERVICE"] = os.Getenv("K_SERVICE")
	k.EnvVars["K_REVISION"] = os.Getenv("K_REVISION")
	k.EnvVars["K_CONFIGURATION"] = os.Getenv("K_CONFIGURATION")

	fsg := &types.FsAccessGroup{
		Group: "FILESYSTEM",
		List: []*types.FsAccessInfo{
			isDirRW("/tmp"),
			isDirRW("/var/log"),
			isDirRW("/dev/log"),
		},
	}
	k.Access = append(k.Access, fsg)

	dnsg := &types.FsAccessGroup{
		Group: "DNS",
		List: []*types.FsAccessInfo{
			isDirRW("/etc/hosts"),
			isDirRW("/etc/hostname"),
			isDirRW("/etc/resolv.conf"),
		},
	}
	k.Access = append(k.Access, dnsg)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(k)

}
