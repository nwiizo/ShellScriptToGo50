package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// ファイルを開く
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 行を読み取りながら、指定のパターンが含まれるかチェック
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "error") {
			fmt.Println(line)
		}
	}
}
