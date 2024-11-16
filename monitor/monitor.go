package monitor

import (
	"Monitoring/memory"
	"time"
)

func Monitor(seconds time.Duration) {
	for {
		memory.PrintMemoryInfo()
		time.Sleep(seconds * time.Second)
	}
}