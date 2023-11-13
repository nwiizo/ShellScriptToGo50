package main

import (
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/path/to/directory", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			os.Chmod(path, 0o444)
		}
		return nil
	})
}
