package main

import "fmt"

func findMinMax(numbers []int) (min int, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	numbers := []int{10, 5, 8, 12, 3, 7}

	min, max := findMinMax(numbers)
	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
}
