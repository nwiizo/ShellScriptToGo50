package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var dirCount, fileCount int

	// Walk through the specified directory
	err := filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", "/path/to/directory", err)
	}

	fmt.Printf("Directories: %d, Files: %d\n", dirCount, fileCount)
}
