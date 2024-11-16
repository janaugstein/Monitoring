package main

import (
	"Monitoring/http"
	"Monitoring/monitor"
	"fmt"
)

func main() {
	fmt.Println("Hello GO")
	go http.ServeUI()
	monitor.Setup()
}