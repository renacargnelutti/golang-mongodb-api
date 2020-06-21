package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"syscall"
)

type OSData struct {
	OSName  string `json:"os_name"`
	DiskUse uint64 `json:"disk_use"`
}

// GetOSData to get data from OS
func GetOSData(w http.ResponseWriter, _ *http.Request) {
	osdata := OSData{
		OSName:  getOSName(),
		DiskUse: getDiskUse(),
	}

	bytes, err := json.Marshal(osdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

// get os name
func getOSName() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" || runtime.GOOS == "FreeBSD" {
		return "Linux"
	} else if runtime.GOOS == "darwin" {
		return "OS X"
	} else {
		return "Undefined"
	}
}

// get disk use
func getDiskUse() uint64 {
	var stat syscall.Statfs_t

	wd, _ := os.Getwd()

	getMemUsage()

	syscall.Statfs(wd, &stat)

	// Available blocks * size per block = available space in bytes
	return (stat.Bavail * uint64(stat.Bsize))
}

// get mem usage
func getMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
