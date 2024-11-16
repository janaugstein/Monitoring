package memory

import (
	"fmt"
	"syscall"
	"time"
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

func GetMemoryStats(seconds time.Duration, msgChannel chan Stats, errChannel chan error) {
	for {
		var memStatus memoryStatusEx
		memStatus.Length = uint32(unsafe.Sizeof(memStatus))

		ret, _, err := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memStatus)))
		if ret == 0 {
			errChannel <- err
			return
		}

		var mem Stats
		mem.Free = float64(memStatus.AvailPhys)
		mem.Total = float64(memStatus.TotalPhys) 
		mem.Used = mem.Total - mem.Free
		mem.PageFileTotal = float64(memStatus.TotalPageFile)
		mem.PageFileFree = float64(memStatus.AvailPageFile)
		mem.VirtualTotal = float64(memStatus.TotalVirtual)
		mem.VirtualFree = float64(memStatus.AvailVirtual)

		msgChannel <- mem

		time.Sleep(seconds * time.Second)
	}
}

func PrintMemoryInfo(memInfo Stats) {
		fmt.Printf("Total Memory: %f GB\n", memInfo.Total/GB)
		fmt.Printf("Free Memory: %f GB\n", memInfo.Free/GB)
		fmt.Printf("Used Memory: %f GB\n", memInfo.Used/GB)	
}

func Test() {
		fmt.Println("msg received")
}

func CheckChannelForMeminfo(channel chan Stats) {
	for {
		memInfo := <- channel
		PrintMemoryInfo(memInfo)
		Test()
	}
}

func CheckErrChannel(errChannel chan error) {
	for {
		err := <-errChannel
		panic(err)
	}
}