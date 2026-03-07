/*
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	Counter atomic.Int64
}

func (c *AtomicCounter) Inc() {
	c.Counter.Add(1)
}

var c AtomicCounter

func main() {

	var wg sync.WaitGroup
	wg.Add(15)
	for _ = range 15 {
		go func() {
			defer wg.Done()
			for _ = range 15 {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Counter is %v", c.Counter.Load())
}
