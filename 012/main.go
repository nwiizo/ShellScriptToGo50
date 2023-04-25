package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := os.Args[1]
	totalSize := int64(0)

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())
			fileInfo, _ := os.Stat(filePath)
			totalSize += fileInfo.Size()
		}
	}

	fmt.Printf("Total size: %d bytes\n", totalSize)
}
