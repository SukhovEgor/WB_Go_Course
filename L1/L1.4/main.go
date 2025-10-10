package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/* Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении. */

func main() {
	// Проверяем, указано ли количество воркеров
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: не введено количество воркеров")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Неверное количество воркеров. Укажите положительное целое число.")
		return
	}

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Канал для обработки SIGINT
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	// Запускаем горутину для обработки сигнала
	go func() {
		<-sigChan
		fmt.Println("\nПолучен сигнал SIGINT, завершаем работу...")
		cancel() // Отменяем контекст при получении SIGINT
	}()

	done := make(chan bool)
	go mainGoroutine(ctx, n, done)

	// Ожидаем завершения главной горутины
	<-done
	fmt.Println("Программа завершена.")
}

func mainGoroutine(ctx context.Context, n int, done chan bool) {
	var wg sync.WaitGroup
	ch := make(chan string, 1)

	// Запускаем N воркеров
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(ctx, ch, i, &wg)
	}

	// Постоянная запись данных в канал
	count := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			wg.Wait() // Ждем завершения всех воркеров
			done <- true
			return
		default:
			time.Sleep(1 * time.Second) // Имитация задержки
			count++
			msg := fmt.Sprintf("Сообщение %d в %v", count, time.Now().Format("15:04:05"))
			select {
			case ch <- msg:
			case <-ctx.Done():
				// Если контекст отменен во время отправки, завершаем
				close(ch)
				wg.Wait()
				done <- true
				return
			}
		}
	}
}

func worker(ctx context.Context, ch <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Если контекст отменен, завершаем воркер
			fmt.Printf("Воркер %d завершает работу по сигналу отмены\n", id)
			return
		case msg, ok := <-ch:
			if !ok {
				// Если канал закрыт, завершаем воркер
				fmt.Printf("Воркер %d завершает работу: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %d получил: %s\n", id, msg)
		}
	}
}
