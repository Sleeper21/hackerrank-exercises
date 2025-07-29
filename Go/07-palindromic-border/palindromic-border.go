package main

import (
	"fmt"
	"strings"
)

// TODO: Yet not solved
// Confusing exercise

func main() {
	input := "ababa"

	output := palindromicBorder(input)

	fmt.Println(output)
}

func palindromicBorder(input string) int32 {
	var prefixList, sufixList []string
	var palindromeBorders int32

	if isPalindrome(input) {
		palindromeBorders++
	}

	for i := 1; i < len(input); i++ {
		prefixList = append(prefixList, input[:i])
		sufixList = append(sufixList, input[i:])
	}

	for _, prefix := range prefixList {
		for _, sufix := range sufixList {
			if prefix == sufix {
				// check if they are palindromes
				if isPalindrome(prefix) {
					palindromeBorders++
					palindromicBorder(prefix)
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}

	fmt.Println(prefixList)
	fmt.Println(sufixList)
	return palindromeBorders
}

func isPalindrome(str string) bool {
	var invertedSlice []string
	var invertedString string

	for i := int32(len(str)) - 1; i >= int32(0); i-- {
		invertedSlice = append(invertedSlice, string(str[i]))
	}

	invertedString = strings.Join(invertedSlice, "")

	if invertedString == str {
		return true
	}

	return false
}
