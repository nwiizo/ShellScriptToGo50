package main

import "fmt"

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	isPalindrome := true
	length := len(input)

	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Printf("%q is a palindrome.\n", input)
	} else {
		fmt.Printf("%q is not a palindrome.\n", input)
	}
}
