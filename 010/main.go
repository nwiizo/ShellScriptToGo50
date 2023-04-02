package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	processID := 12345

	process, err := os.FindProcess(processID)
	if err != nil {
		fmt.Printf("Error finding process with ID %d: %v\n", processID, err)
		return
	}

	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Printf("Error killing process with ID %d: %v\n", processID, err)
	} else {
		fmt.Printf("Process with ID %d killed successfully\n", processID)
	}
}
