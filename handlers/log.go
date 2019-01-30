package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	maxLogSizeInKB = int64(1024) //1MB
	// TODO: parameterize this like maxLogSizeInKB
	supportedLogPaths = []string{"/tmp", "/var/log", "/dev/log"}
)

func listDirContent(w http.ResponseWriter, path string) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Fprintf(w, "   %dKB - %s\n", getSizeInKB(f.Size()), f.Name())
	}

	return nil
}

func getSizeInKB(size int64) int64 {
	if size < 1 {
		return 0
	}
	fileSize := size / 1024 // in kilobytes
	return fileSize
}

func logHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Log...")

	// gt log path
	path := r.URL.Query().Get("logpath")
	path = filepath.Clean(path)

	if path == "" {
		fmt.Fprintf(w, "The `logpath` parameter required (e.g. /log?logpath=/var/log/ktest.log)\n")
		fmt.Fprintf(w, "Supported log locations:\n")
		for _, p := range supportedLogPaths {
			fmt.Fprintf(w, "   %s/*\n", p)
		}
		return
	}
	log.Printf("logpath: %s", path)

	// check if one of the allowed paths
	ifValidLogDir := false
	for _, p := range supportedLogPaths {
		if filepath.HasPrefix(path, p) {
			ifValidLogDir = true
			break
		}
	}

	// if not a valid path root
	if !ifValidLogDir {
		fmt.Fprintf(w, "Invalid log file path: %s\n", path)
		fmt.Fprintf(w, "Supported log locations:\n")
		for _, p := range supportedLogPaths {
			fmt.Fprintf(w, "   %s/*\n", p)
		}
		return
	}

	// open log file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(w, "404: Log not found: %s - %v", path, err)
		return
	}

	// header
	fmt.Fprintf(w, "============================================================\n")
	fmt.Fprintf(w, "Content: %s\n", path)
	fmt.Fprintf(w, "============================================================\n")

	// stats
	fileStats, _ := file.Stat()

	// dir
	if fileStats.IsDir() {
		listDirContent(w, path)
		return
	}

	// file
	fileSizeKB := getSizeInKB(fileStats.Size())
	log.Printf("Log size: %dKB", fileSizeKB)

	// check if max log size was set in env vars
	maxLogSizeParam := os.Getenv("MAX_LOG_SIZE_IN_KB")
	if maxLogSizeParam != "" {
		maxLogSizeInKB, err = strconv.ParseInt(maxLogSizeParam, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Invalid MAX_LOG_SIZE_IN_KB parameter: %s", maxLogSizeParam)
			return
		}
	}

	// check if too large
	if fileSizeKB > maxLogSizeInKB {
		fmt.Fprintf(w, "Log size (%dKB) exceeds max allowed size (%dKB)", fileSizeKB, maxLogSizeInKB)
		return
	}

	// copy content to the response writer
	io.Copy(w, file)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	return

}
