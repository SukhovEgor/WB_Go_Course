/*
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент,
возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, -1, 0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Position of %d is %d in arr %v", 2, binarySearch(nums, 2), nums)

}

func binarySearch(nums []int, target int) int {
	min, max := 0, len(nums)-1
	for min < max {
		mid := (min + max) / 2
		switch {
		case target == nums[mid]:
			return mid
		case target < nums[mid]:
			max = mid - 1
		case target > nums[mid]:
			min = mid + 1
		}
	}
	return -1
}
