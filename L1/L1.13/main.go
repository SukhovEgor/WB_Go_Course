package main

import (
	"fmt"
	"math/rand"
)

/* Поменять местами два числа без использования временной переменной.
Подсказка: примените сложение/вычитание или XOR-обмен. */

func main() {
	a := rand.Intn(1000)
	b := rand.Intn(1000)
	fmt.Println("Before swap:", a, b)

	xorSwap(a, b)
	arithmeticSwap(a, b)
	standardSwap(a, b)
}

func xorSwap(a int, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("XOR swap:", a, b)
}
func arithmeticSwap(a int, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Arithmetic swap:", a, b)
}
func standardSwap(a int, b int) {
	a, b = b, a
	fmt.Println("Standard swap:", a, b)
}
