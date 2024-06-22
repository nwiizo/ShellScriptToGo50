package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	logFile, err := os.OpenFile("/var/log/container_monitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	for {
		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
		if err != nil {
			log.Printf("Error listing containers: %v", err)
			continue
		}

		for _, container := range containers {
			if container.State != "running" {
				restartContainer(cli, container.ID)
			} else {
				logResourceUsage(cli, container.ID)
			}
		}

		time.Sleep(60 * time.Second)
	}
}

func restartContainer(cli *client.Client, containerID string) {
	log.Printf("Restarting container %s", containerID)
	if err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
		log.Printf("Error restarting container %s: %v", containerID, err)
	}
}

func logResourceUsage(cli *client.Client, containerID string) {
	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		log.Printf("Error getting stats for container %s: %v", containerID, err)
		return
	}
	defer stats.Body.Close()

	var statsJSON types.StatsJSON
	if err := json.NewDecoder(stats.Body).Decode(&statsJSON); err != nil {
		log.Printf("Error decoding stats for container %s: %v", containerID, err)
		return
	}

	cpuPercent := calculateCPUPercentUnix(statsJSON.PreCPUStats.CPUUsage.TotalUsage, statsJSON.CPUStats.CPUUsage.TotalUsage)
	memoryUsage := float64(statsJSON.MemoryStats.Usage) / 1024 / 1024 // Convert to MB

	log.Printf("Container %s - CPU: %.2f%%, Memory: %.2f MB", containerID, cpuPercent, memoryUsage)
}

func calculateCPUPercentUnix(previousCPU, currentCPU uint64) float64 {
	cpuDelta := float64(currentCPU - previousCPU)
	return (cpuDelta / float64(100*time.Second/time.Nanosecond)) * 100.0
}
