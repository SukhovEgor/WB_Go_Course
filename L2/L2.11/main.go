//Напишите функцию, которая находит все множества анаграмм по заданному словарю.

package main

import (
	"fmt"
	"slices"
	"strings"
)


func GetAnagrams(words []string) map[string][]string {
	if len(words) == 0 {
		return nil
	}

	groups := make(map[string][]string)
	firstSeen := make(map[string]string)

	for _, word := range words {
		if word == "" {
			continue
		}

		lower := strings.ToLower(word)

		runes := []rune(lower)
		slices.Sort(runes)
		key := string(runes)

		if _, exists := firstSeen[key]; !exists {
			firstSeen[key] = word
		}

		groups[key] = append(groups[key], word)
	}

	result := make(map[string][]string, len(groups))

	for key, group := range groups {
		if len(group) <= 1 {
			continue
		}

		slices.Sort(group)
		result[firstSeen[key]] = group
	}

	return result
}

func main() {
	input := []string{
		"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик",
		"стол", "ПЯТКА", "Тяпка",
	}

	for key, group := range GetAnagrams(input) {
		fmt.Printf("%s : %v\n", key, group)
	}
}