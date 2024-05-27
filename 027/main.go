package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

func main() {
	directory := ".."

	size, err := getDirSize(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Directory size: %d bytes\n", size)
}
