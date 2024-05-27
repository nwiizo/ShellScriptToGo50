package main

import "fmt"

func main() {
	var numbers [5]int

	fmt.Println("Enter 5 numbers:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Average: %.2f\n", average)
}
