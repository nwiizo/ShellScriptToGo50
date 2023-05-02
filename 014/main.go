package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Size int64
}

func main() {
	targetDir := os.Args[1]

	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			info := FileInfo{
				Path: filepath.Join(targetDir, file.Name()),
				Size: file.Size(),
			}
			fmt.Printf("File: %s, Size: %d bytes\n", info.Path, info.Size)
		}
	}
}
