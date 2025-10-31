package main

import (
	"fmt"
	"os"
	"strconv"
)

/* Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
Подсказка: используйте битовые операции (|, &^). */

// Заменяет i-й бит в 1 или 0
func changeBit(num int64, i int, bit int) (int64, error) {
	if i < 0 || i > 63 {
		return 0, fmt.Errorf("Ошибка: бит %d вне диапазона! Доступно 0–63", i)
	}
	if bit != 0 && bit != 1 {
		return 0, fmt.Errorf("Jшибка: значение бита должно быть 0 или 1, а не %d", bit)
	}

	if bit == 1 {
		return num | (1 << i), nil
	}
	return num &^ (1 << i), nil
}

func printBinary(num int64) {
	fmt.Print("Двоичный вид: ")
	for i := 63; i >= 0; i-- {
		if (num>>i)&1 == 1 {
			fmt.Print("1")
			// После первой 1 печатаем все остальные биты
			for j := i - 1; j >= 0; j-- {
				if (num>>j)&1 == 1 {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
			fmt.Println()
			return
		}
	}
}

func main() {
	fmt.Println("Установим бит в число!")

	if len(os.Args) != 4 {
		fmt.Println("Неверный формат ввода: go run main.go <число> <позиция_бита> <0_или_1>")
		return
	}

	numStr := os.Args[1]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Printf("Введите целое число!")
		return
	}

	iStr := os.Args[2]
	i, err := strconv.Atoi(iStr)
	if err != nil || i < 0 {
		fmt.Printf("Неверный формат ввода позиции: Должно быть 0–63!")
		return
	}

	bitStr := os.Args[3]
	bit, err := strconv.Atoi(bitStr)
	if err != nil {
		fmt.Printf("Неверный формат ввода бита: Должно быть  0 или 1!")
		return
	}

	fmt.Printf("Было: %d\n", num)
	printBinary(num)

	// Замена бита
	result, err := changeBit(num, i, bit)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("\nЗамена %d-ого бита в %d...\n", i, bit)
	fmt.Printf("Стало: %d\n", result)
	printBinary(result)
}
