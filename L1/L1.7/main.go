package main

import (
	"fmt"
	"sync"
)

/* Реализовать безопасную для конкуренции запись данных в структуру map.
Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
Проверьте работу кода на гонки (util go run -race). */

type SafeMap struct {
	data map[string]int
	mu   sync.Mutex 
}

// Запись значения в map
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	sm.data[key] = value
	sm.mu.Unlock()   
}

// Чтение значения
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, exists := sm.data[key]
	return value, exists
}

func main() {
	sm := SafeMap{
		data: make(map[string]int),
	}

	var wg sync.WaitGroup

	// Запускаем 10 горутин, которые пишут в одну и ту же карту
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() 

			key := fmt.Sprintf("воркер-%d", id)
			value := id * 100

			sm.Set(key, value)
			fmt.Printf("Горутина %d записала: %s = %d\n", id, key, value)
		}(i)
	}

	wg.Wait()

	fmt.Println("\nИтоговый map:")
	for key, value := range sm.data {
		fmt.Printf("  %s: %d\n", key, value)
	}

	fmt.Println("\nГотово.")
}