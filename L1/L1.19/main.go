/* Разработать программу, которая переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
 то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune). */

package main

import (
	"fmt"
)

func main() {
	str := "Главрыба"
	fmt.Printf("Reverse of %s is %s", str, Reverse(str))

}

func Reverse(str string) string {
	letters := []rune(str)
	for i := 0; i < len(letters)/2; i++ {
		letters[i], letters[len(letters)-i-1] = letters[len(letters)-i-1], letters[i]

	}
	return string(letters)
}
