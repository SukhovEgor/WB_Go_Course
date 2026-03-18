package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("Is string '%s' unique? %t\n", "abcd", UniqueChars("abcd"))
	fmt.Printf("Is string '%s' unique? %t\n", "EWRfgdwer", UniqueChars("EWRfgdwer"))
	fmt.Printf("Is string '%s' unique? %t\n", "Uniq1e", UniqueChars("Uniq1e"))
	fmt.Printf("Is string '%s' unique? %t\n", "12416", UniqueChars("12416"))
	fmt.Printf("Is string '%s' unique? %t\n", "Hello world!", UniqueChars("Hello world!"))
}

func UniqueChars(str string) bool {
	charMap := make(map[rune]bool)

	for _, char := range str {
		lowerChar := unicode.ToLower(char)

		if _, ok := charMap[lowerChar]; ok {
			return false
		}
		charMap[lowerChar] = true
	}
	return true

}
