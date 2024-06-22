package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sourceDir := "/path/to/source"

	// サブディレクトリを作成
	subdirs := []string{"documents", "images", "videos", "others"}
	for _, dir := range subdirs {
		os.MkdirAll(filepath.Join(sourceDir, dir), os.ModePerm)
	}

	// ファイルを整理
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		var destDir string

		switch ext {
		case ".pdf", ".doc", ".docx", ".txt":
			destDir = "documents"
		case ".jpg", ".jpeg", ".png", ".gif":
			destDir = "images"
		case ".mp4", ".avi", ".mov":
			destDir = "videos"
		default:
			destDir = "others"
		}

		oldPath := filepath.Join(sourceDir, file.Name())
		newPath := filepath.Join(sourceDir, destDir, file.Name())
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error moving file %s: %v\n", file.Name(), err)
		}
	}

	fmt.Println("File organization complete")
}
