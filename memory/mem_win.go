package memory

import (
	"fmt"
	"syscall"
	"unsafe"
)


type Stats struct {
	Total, Used, Free, PageFileTotal, PageFileFree, VirtualTotal, VirtualFree float64
}

type memoryStatusEx struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var globalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
var KB float64 = 1028
var MB float64 = KB * 1028
var GB float64 = MB * 1028
	

func GetMemoryStats() (*Stats, error) {
	var memStatus memoryStatusEx
	memStatus.Length = uint32(unsafe.Sizeof(memStatus))

	ret, _, err := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memStatus)))
	if ret == 0 {
		return nil, fmt.Errorf("failed in globalMemoryStatusEx: %s", err)
	}

	var mem Stats
	mem.Free = float64(memStatus.AvailPhys)
	mem.Total = float64(memStatus.TotalPhys) 
	mem.Used = mem.Total - mem.Free
	mem.PageFileTotal = float64(memStatus.TotalPageFile)
	mem.PageFileFree = float64(memStatus.AvailPageFile)
	mem.VirtualTotal = float64(memStatus.TotalVirtual)
	mem.VirtualFree = float64(memStatus.AvailVirtual)

	return &mem, nil
}

func PrintMemoryInfo() {
	memInfo, err := GetMemoryStats()
	if err != nil {
		panic(err)
	}


	fmt.Printf("Total Memory: %f GB\n", memInfo.Total/GB)
	fmt.Printf("Free Memory: %f GB\n", memInfo.Free/GB)
	fmt.Printf("Used Memory: %f GB\n", memInfo.Used/GB)
}