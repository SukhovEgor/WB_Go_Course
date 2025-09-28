package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, value := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(value)
	}
	wg.Wait()
}
