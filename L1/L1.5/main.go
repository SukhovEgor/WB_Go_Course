package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)
/* Разработать программу, которая будет последовательно отправлять значения в канал,
 а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться. */

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: укажите количество секунд")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: укажите положительное число секунд")
		return
	}


	ch := make(chan string)

	// Запускаем горутину для чтения
	go func () {
	for msg := range ch {
		fmt.Printf("Получено: %s\n", msg)
	}
	fmt.Println("Чтение завершено: канал закрыт")
}()

	// Отправляем сообщения до истечения времени
	timer := time.After(time.Duration(n) * time.Second)
	count := 0
	for {
		select {
		case <-timer:
			// Время истекло, закрываем канал и завершаем
			close(ch)
			fmt.Println("Программа завершена.")
			return
		default:
			count++
			ch <- fmt.Sprintf("Сообщение %d в %v", count, time.Now().Format("15:04:05"))
			time.Sleep(1 * time.Second) // Задержка для имитации
		}
	}
}

