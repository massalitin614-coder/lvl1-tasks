package main

import (
	"fmt"
	"sync"
)

// Counter - потокобезопасный счётчик
type Counter struct {
	mu    sync.RWMutex
	value int
}

// Increment увеличивает значение на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Add увеличивает значение на n (полезно для пакетных операций)
func (c *Counter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += n
}

// Value возвращает текущее значение
func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// Reset сбрасывает счётчик в 0
func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Конфигурация
	numWorkers := 100
	incrementsPerWorker := 100

	fmt.Printf("Запускаем %d горутин, по %d инкрементов каждая\n",
		numWorkers, incrementsPerWorker)

	// Запуск горутин
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			// Можно использовать Add для оптимизации:
			// counter.Add(incrementsPerWorker)
			// Но если нужно симулировать реальную работу, делаем цикл:
			for j := 0; j < incrementsPerWorker; j++ {
				counter.Increment()
			}

			if workerID%10 == 0 {
				fmt.Printf("Горутина %d завершила работу\n", workerID)
			}
		}(i)
	}

	wg.Wait()

	// Проверяем результат
	expected := numWorkers * incrementsPerWorker
	actual := counter.Value()

	fmt.Printf("\n=== Результаты ===\n")
	fmt.Printf("Итоговое значение: %d\n", actual)
	fmt.Printf("Ожидаемое значение: %d\n", expected)

	if actual == expected {
		fmt.Println("Счётчик работает корректно!")
	} else {
		fmt.Printf("Ошибка! Расхождение в %d\n", expected-actual)
	}

	// Демонстрация других методов
	fmt.Println("\n=== Дополнительные тесты ===")
	counter.Reset()
	fmt.Printf("После Reset: %d\n", counter.Value())

	counter.Add(50)
	fmt.Printf("После Add(50): %d\n", counter.Value())
}
