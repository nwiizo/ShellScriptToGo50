package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	processName := "docker"

	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, processName) {
			fmt.Println(line)
		}
	}
}
