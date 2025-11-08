package main

import (
	"fmt"
)

/* Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива,
во второй – результат операции x*2. После этого данные из второго канала должны выводиться
в stdout. То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их
 обработка. Убедитесь, что чтение из второго канала корректно завершается. */

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

func print(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	src := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	doubled := double(src)
	print(doubled)
}
