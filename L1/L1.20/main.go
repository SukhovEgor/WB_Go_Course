/* Разработать программу, которая переворачивает порядок слов в строке.
Пример: входная строка:
«snow dog sun», выход: «sun dog snow».
Считайте, что слова разделяются одиночным пробелом. Постарайтесь
не использовать дополнительные срезы, а выполнять операцию «на месте». */

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sun dog snow moon 1"
	fmt.Printf("Reverse of %s is %s", str, Reverse(str))
}

func Reverse(str string) string {
	words := strings.Fields(str)
	start, end := 0, len(words) - 1
	for start <= end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
	return strings.Join(words, " ")
}