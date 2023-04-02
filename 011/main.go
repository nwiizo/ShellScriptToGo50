package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// backupScript represents your backup script function.
func backupScript() {
	fmt.Println("Backup script executed:", time.Now())
}

func main() {
	c := cron.New(cron.WithSeconds()) // Use cron.WithSeconds() to support second-level granularity
	_, err := c.AddFunc("0 0 3 * * *", backupScript)
	if err != nil {
		fmt.Println("Error scheduling the backup script:", err)
		return
	}

	c.Start()
	defer c.Stop()

	// Keep the main function running
	select {}
}
