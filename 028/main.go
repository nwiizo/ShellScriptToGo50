package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func syncDirectories(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			os.MkdirAll(dstPath, info.Mode())
		} else {
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func main() {
	srcDir := "/path/to/source"
	dstDir := "/path/to/destination"

	err := syncDirectories(srcDir, dstDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Directories synchronized successfully.")
}
