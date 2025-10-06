package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type SystemResource struct {
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
}

func getSystemResource() SystemResource {
	cpuUsage, _ := cpu.Percent(0, false)
	memoryUsage, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")

	return SystemResource{
		CPUUsage:    cpuUsage[0],
		MemoryUsage: memoryUsage.UsedPercent,
		DiskUsage:   diskUsage.UsedPercent,
	}
}

func moveTo(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}
func refreshSystemResource() {
	for {
		resource := getSystemResource()
		clearScreen()
		moveTo(10, 5)
		fmt.Printf("CPU Usages: %.2f%%\n", resource.CPUUsage)
		moveTo(10, 6)
		fmt.Printf("Memory Usages: %.2f%%\n", resource.MemoryUsage)
		moveTo(10, 7)
		fmt.Printf("Disk Usages: %.2f%%\n", resource.DiskUsage)

		time.Sleep(3 * time.Second)
	}
}

func SystemMonitor() {
	go refreshSystemResource()
	// to avoid main program jammed
	select {} // make sure main program don't exit immediately
}
