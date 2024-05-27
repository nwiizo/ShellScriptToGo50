package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "sample.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %d characters\n", lineNumber, len(line))
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
