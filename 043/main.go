package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	logFile := "/var/log/apache2/access.log"

	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	ipCounts := make(map[string]int)
	urlCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 6 {
			ipCounts[fields[0]]++
			urlCounts[fields[6]]++
		}
	}

	fmt.Println("Top 5 IP addresses:")
	printTopN(ipCounts, 5)

	fmt.Println("\nTop 5 requested URLs:")
	printTopN(urlCounts, 5)
}

func printTopN(counts map[string]int, n int) {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i := 0; i < n && i < len(ss); i++ {
		fmt.Printf("%s: %d\n", ss[i].Key, ss[i].Value)
	}
}
