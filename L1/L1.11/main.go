package main

import (
	"fmt"
	"sort"
)

func intersection(a, b []int) []int {
	seen := make(map[int]bool)

	for _, v := range a {
		seen[v] = true
	}

	var result []int
	for _, v := range b {
		if seen[v] {
			result = append(result, v)
			delete(seen, v)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	A := []int{1, 2, 3, 5, 7}
	B := []int{2, 3, 4, 6, 7}

	fmt.Printf("%v\n", intersection(A, B))
}
