package monitor

import (
	"Monitoring/memory"
)

var memoryChannel = make(chan memory.Stats)
var memErrorChannel = make(chan error)

func Setup() {
	go memory.CheckChannelForMeminfo(memoryChannel)
	go memory.CheckErrChannel(memErrorChannel)
	memory.GetMemoryStats(2, memoryChannel, memErrorChannel)
}