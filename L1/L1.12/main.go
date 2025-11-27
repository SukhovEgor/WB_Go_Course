package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "cat", "dog","cat", "tree","dog", "cat", "tree"}
	seen := make(map[string]struct{})
	for _, word := range words {
		seen[word] = struct{}{}
	}

	result := make([]string, 0, len(seen))
	for word := range seen {
		result = append(result, word)
	}

	sort.Strings(result)
	fmt.Println(result)
}
