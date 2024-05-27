package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	cmd := exec.Command("ls", "-l")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("output_%s.txt", timestamp)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Command output saved to %s\n", filename)

	files, err := filepath.Glob("output_*.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saved files:")
	for _, file := range files {
		fmt.Println(file)
	}
}
