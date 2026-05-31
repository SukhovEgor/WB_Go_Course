/* Написать функцию Go, осуществляющую примитивную распаковку строки,
содержащей повторяющиеся символы/руны. */

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	result.Grow(len(s) * 2)

	runes := []rune(s)
	i := 0

	for i < len(runes) {
		curr := runes[i]

		// Обработка экранирования
		if curr == '\\' && i+1 < len(runes) {
			i++
			result.WriteRune(runes[i])
			i++
			continue
		}

		// Если это цифра, то повторяем предыдущий символ
		if unicode.IsDigit(curr) {
			if result.Len() == 0 {
				return "", ErrInvalidString
			}

			numStr := string(curr)
			i++
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				numStr += string(runes[i])
				i++
			}

			count, err := strconv.Atoi(numStr)
			if err != nil || count < 0 {
				return "", ErrInvalidString
			}

			// Повторяем последний символ count раз
			lastChar := result.String()[result.Len()-1:]
			result.WriteString(strings.Repeat(lastChar, count-1)) 
			continue
		}

		// Обычный символ
		result.WriteRune(curr)
		i++
	}

	return result.String(), nil
}

func main() {
	testCases := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
	}

	for _, input := range testCases {
		result, err := Unpack(input)
		if err != nil {
			fmt.Printf("Input: %q | Error: %v\n", input, err)
		} else {
			fmt.Printf("Input: %q | Output: %q\n", input, result)
		}
	}
}
