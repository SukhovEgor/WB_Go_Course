package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// Закрытие канала
func stopByChannelClose() {
	fmt.Println("\nЗакрытие канала")
	ch := make(chan string)

	go func() {
		for msg := range ch {
			fmt.Printf("Горутина получила: %s\n", msg)
		}
		fmt.Println("Горутина завершена: канал закрыт")
	}()

	ch <- "Сообщение 1"
	time.Sleep(200 * time.Millisecond)
	ch <- "Сообщение 2"
	time.Sleep(200 * time.Millisecond)

	close(ch)
	time.Sleep(500 * time.Millisecond) 
}

// Контекст с отменой
func stopByContext() {
	fmt.Println("\nКонтекст с отменой")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена: контекст отменён")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Отменяем контекст")
	cancel()
	time.Sleep(500 * time.Millisecond) 
}

// Атомарный флаг
func stopByAtomicFlag() {
	fmt.Println("\nАтомарный флаг")
	var running int32 = 1

	go func() {
		for atomic.LoadInt32(&running) == 1 {
			fmt.Println("Горутина работает...")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Горутина завершена: флаг сброшен")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Сбрасываем флаг")
	atomic.StoreInt32(&running, 0)
	time.Sleep(500 * time.Millisecond)
}

// Таймер через time.After
func stopByTimer() {
	fmt.Println("\nТаймер через time.After")
	go func() {
		timer := time.After(2 * time.Second)
		for {
			select {
			case <-timer:
				fmt.Println("Горутина завершена: таймер сработал")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// Ждём, пока таймер сработает
	time.Sleep(3 * time.Second)
}

// Канал уведомления (done)
func stopByDoneChannel() {
	fmt.Println("\nКанал уведомления (done)")
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина завершена: сигнал через done")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Отправляем сигнал done")
	close(done)
	time.Sleep(500 * time.Millisecond)
}

// runtime.Goexit
// Горутина завершается принудительно через runtime.Goexit
func stopByGoexit() {
	fmt.Println("\nruntime.Goexit")
	go func() {
		fmt.Println("Горутина работает...")
		time.Sleep(1 * time.Second)
		fmt.Println("Вызываем Goexit")
		runtime.Goexit()
		fmt.Println("Это не выведется, так как Goexit завершил горутину")
	}()

	// Ждём завершения
	time.Sleep(2 * time.Second)
	fmt.Println("Горутина завершена через Goexit")
}

func main() {

	stopByChannelClose()
	stopByContext()
	stopByAtomicFlag()
	stopByTimer()
	stopByDoneChannel()
	stopByGoexit()

	fmt.Println("\nВсе примеры завершены!")
}
