/* Разработать программу, которая перемножает, делит, складывает, вычитает две числовых 
переменных a, b, значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами, но обратите 
внимание на возможное переполнение для ещё больших значений.
Для очень больших чисел можно использовать math/big. */

package main

import (
	"fmt"
	"math/rand"
	"math/big"

)

func main() {
	numA := generateRandNum()
	numB := generateRandNum()

	fmt.Println("a = ", numA)
	fmt.Println("b = ", numB)

	mathOperations := new(big.Float)

	fmt.Println("a + b = ",mathOperations.Add(numA,numB))
	fmt.Println("a - b = ",mathOperations.Sub(numA,numB))
	fmt.Println("a * b = ",mathOperations.Mul(numA,numB))
	fmt.Println("a / b = ",mathOperations.Quo(numA,numB))
}


func generateRandNum() *big.Float {
	min := int64(1 << 40)
	max := int64(1 << 50)

	return big.NewFloat(float64(min + rand.Int63n(max - min + 1)))
}