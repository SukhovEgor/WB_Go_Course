/*
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
(copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSlice := generateSlice(5, 100)
	fmt.Println("Original slice is", intSlice)

	i := rand.Intn(len(intSlice))
	fmt.Println("Remove element", intSlice[i], "at index", i)

	result := removeElement(intSlice, i)
	fmt.Println("Result is", result)
}

func generateSlice(minValue int, maxValue int) (result []int) {
	for i := 0; i <= 10; i++ {
		result = append(result, (minValue + rand.Intn(maxValue - minValue + 1)))
	}
	return result
}

func removeElement(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		return arr
	}

	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = 0

	return arr[:len(arr)-1]
}