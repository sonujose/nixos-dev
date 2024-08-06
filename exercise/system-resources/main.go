package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func main() {
	for {
		printCPUUsage()
		printMemoryUsage()
		printDiskUsage()
		printNetworkUsage()

		// Sleep for a specified interval before the next update
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

func printMemoryUsage() {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting memory usage: %v\n", err)
		return
	}

	fmt.Printf("Memory Usage:\n")
	fmt.Printf("  Total: %.2f GB\n", float64(v.Total)/1024/1024/1024)
	fmt.Printf("  Used: %.2f GB\n", float64(v.Used)/1024/1024/1024)
	fmt.Printf("  Free: %.2f GB\n", float64(v.Free)/1024/1024/1024)
	fmt.Printf("  UsedPercent: %.2f%%\n", v.UsedPercent)
}

func printDiskUsage() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("Error getting disk partitions: %v\n", err)
		return
	}

	fmt.Println("Disk Usage:")
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			fmt.Printf("  Error getting usage for %s: %v\n", p.Mountpoint, err)
			continue
		}

		fmt.Printf("  %s:\n", p.Mountpoint)
		fmt.Printf("    Total: %.2f GB\n", float64(usage.Total)/1024/1024/1024)
		fmt.Printf("    Used: %.2f GB\n", float64(usage.Used)/1024/1024/1024)
		fmt.Printf("    Free: %.2f GB\n", float64(usage.Free)/1024/1024/1024)
		fmt.Printf("    UsedPercent: %.2f%%\n", usage.UsedPercent)
	}
}

func printNetworkUsage() {
	counters, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("Error getting network usage: %v\n", err)
		return
	}

	fmt.Println("Network Usage:")
	for _, counter := range counters {
		fmt.Printf("  %s:\n", counter.Name)
		fmt.Printf("    Bytes Sent: %v\n", counter.BytesSent)
		fmt.Printf("    Bytes Received: %v\n", counter.BytesRecv)
		fmt.Printf("    Packets Sent: %v\n", counter.PacketsSent)
		fmt.Printf("    Packets Received: %v\n", counter.PacketsRecv)
	}
}
