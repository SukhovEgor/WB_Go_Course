package main

import (
	"fmt"
)

/* Разработать программу, которая в runtime способна определить тип переменной,
переданной в неё (на вход подаётся interface{}).
 Типы, которые нужно распознавать: int, string, bool, chan (канал).
Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func TypePrint(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("%s is String type\n", v)
	case int:
		fmt.Printf("%d is Integer type\n", v)
	case bool:
		fmt.Printf("%v is Boolean type\n", v)
	case chan int:
		fmt.Printf("%v is Channel int type\n", v)
	}
}

func main() {
	str := "Hello!"
	var num int
	var bool bool
	ch := make(chan int)
	TypePrint(str)
	TypePrint(num)
	TypePrint(bool)
	TypePrint(ch)
}
