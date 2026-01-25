package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap - потокобезопасная map
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap создает новую потокобезопасную map
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set устанавливает значение по ключу
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get возвращает значение по ключу и флаг существования
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	value, exists := sm.data[key]
	if !exists {
		return 0, false
	}
	return value, true
}

// Delete удаляет значение по ключу
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Increment увеличивает значение по ключу на 1
func (sm *SafeMap) Increment(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key]++
}

// GetAll возвращает копию всех данных
func (sm *SafeMap) GetAll() map[string]int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	result := make(map[string]int, len(sm.data))
	for key, value := range sm.data {
		result[key] = value
	}
	return result
}

// Len возвращает количество элементов в map
func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// Clear очищает map
func (sm *SafeMap) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data = make(map[string]int)
}

func main() {
	// Создаем потокобезопасную map
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	// Запускаем 10 горутин для записи
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("goroutine-%d", id)
				safeMap.Increment(key)
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	// Запускаем горутину для чтения
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(10 * time.Millisecond)
			length := safeMap.Len()
			fmt.Printf("Чтение %d: найдено %d записей\n", i+1, length)
		}
	}()

	// Ждем завершения всех горутин
	wg.Wait()

	fmt.Println("\nФинальный результат:")
	for key, value := range safeMap.GetAll() {
		fmt.Printf("%s: %d\n", key, value)
	}
}
