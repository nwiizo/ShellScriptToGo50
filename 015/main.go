package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path      string
	Extension string
}

type FileLister interface {
	ListFiles(dir string) ([]FileInfo, error)
}

type LocalFileLister struct{}

func (l LocalFileLister) ListFiles(dir string) ([]FileInfo, error) {
	var fileInfos []FileInfo
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == "" {
				ext = "other"
			}
			fileInfos = append(fileInfos, FileInfo{
				Path:      path,
				Extension: ext,
			})
		}
		return nil
	})

	return fileInfos, err
}

func main() {
	targetDir := os.Args[1]

	lister := LocalFileLister{}
	fileInfos, err := lister.ListFiles(targetDir)
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}

	extensionCount := make(map[string]int)
	for _, fileInfo := range fileInfos {
		extensionCount[fileInfo.Extension]++
	}

	for ext, count := range extensionCount {
		fmt.Printf("Extension: %s, Count: %d\n", ext, count)
	}
}
