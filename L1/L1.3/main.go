package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

/* Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров. */

func main() {
	done := make(chan bool)
	
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: не введено количество воркеров")
		return
	} 

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Неверное количество воркеров. Укажите положительное целое число.")
		return
	}

	go mainGoroutine(n, done)
	<- done

}

func mainGoroutine(n int, done chan bool) {
	var wg sync.WaitGroup
	ch := make(chan string, 1) 	
	// Запускаем N воркеров с использованием WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(ch, i, &wg)
	}

	// Постоянная запись данных в канал в главной горутине (для демонстрации используем 10 сообщений)
	count := 0
	for count < 10 {
		time.Sleep(1 * time.Second) // Имитация задержки
		count++
		msg := fmt.Sprintf("Сообщение %d в %v", count, time.Now().Format("15:04:05"))
		ch <- msg
	}

	close(ch)
	done <- true
	// Ждем завершения всех воркеров
	wg.Wait()
	fmt.Println("Все воркеры завершили работу.")
}

func worker(ch <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("Воркер %d получил: %s\n", id, msg)
	}
}