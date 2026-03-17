/*
Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
которая приостанавливает выполнение текущей горутины.
Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение
(например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.

Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	sleep := 10 * time.Second
	fmt.Println("Ожидание:", sleep)

	fmt.Println("Канал с таймером начался")
	SleepChan(sleep)
	fmt.Println("Канал с таймером закончился")
	fmt.Println()
	fmt.Println("Простой канал time.After начался")
	SleepTimer(sleep)
	fmt.Println("Простой канал time.After закончился")
	fmt.Println()
	fmt.Println("Ожидание через контекст началось")
	SleepContext(sleep)
	fmt.Println("Ожидание через контекст закончилось")
}

func SleepChan(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C //Блокается пока не придет сигнал
}

func SleepTimer(duration time.Duration) {
	<-time.After(duration)
}

func SleepContext(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	<-ctx.Done()
}