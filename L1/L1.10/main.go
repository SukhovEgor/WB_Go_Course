package main

import (
	"fmt"
	"sort"
)

func toDecade(n float64) int {
	return int(n) / 10 * 10
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	// Группируем
	for _, t := range temps {
		key := toDecade(t)
		groups[key] = append(groups[key], t)
	}

	// Сортируем ключи
	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		values := groups[k]
		// Сортируем внутри группы
		sort.Float64s(values)
		fmt.Printf("%d:{", k)
		for i, v := range values {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", v)
		}
		fmt.Println("}")
	}
}