package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	s1 := []rune(str1)
	s2 := []rune(str2)

	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	return string(s1) == string(s2)
}

func main() {
	str1 := "listen"
	str2 := "silent"

	if isAnagram(str1, str2) {
		fmt.Printf("%q and %q are anagrams.\n", str1, str2)
	} else {
		fmt.Printf("%q and %q are not anagrams.\n", str1, str2)
	}
}
