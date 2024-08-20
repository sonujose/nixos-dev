package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	for {
		printCPUUsage()
		time.Sleep(5 * time.Second)
	}
}

func printCPUUsage() {
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Printf("Error getting CPU usage: %v\n", err)
		return
	}

	fmt.Println("CPU Usage:")
	for i, percentage := range percentages {
		fmt.Printf("  CPU%d: %.2f%%\n", i, percentage)
	}
}
