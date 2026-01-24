package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// 1. Выход по условию (Condition-based exit)
// Простейший способ - горутина завершается при достижении внутреннего условия
// Плюсы: простота, полный контроль изнутри
// Минусы: нет возможности внешнего управления
func conditionExit(wg *sync.WaitGroup) {
	defer wg.Done() // Всегда вызываем Done при завершении, defer гарантирует выполнение

	count := 0
	for count < 3 { // Условие завершения прямо в цикле - самый простой подход
		fmt.Printf("Условие: итерация %d\n", count)
		time.Sleep(300 * time.Millisecond)
		count++
	}
	fmt.Println("Завершение по условию")
}

// 2. Через канал уведомления (Notification channel)
// Классический Go-паттерн: используем канал для сигнализации о завершении
// Плюсы: гибкость, возможность уведомить несколько горутин
// Минусы: нужно управлять каналом извне
func channelExit(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Бесконечный цикл с проверкой сигнала остановки
	for {
		select {
		case <-stopChan: // Ждем сигнал остановки (закрытие канала или отправка значения)
			fmt.Println("Получен сигнал из канала")
			return // Завершаем горутину
		default:
			fmt.Println("Работаю")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// 3. Через контекст (Context)
// Современный идиоматичный способ, особенно для HTTP/gRPC серверов
// Плюсы: стандартизированный API, цепочки отмен, таймауты
// Минусы: немногословный (меньше информации о причинах отмены)
func contextExit(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): // Контекст отменен (cancel() вызван, таймаут или дедлайн)
			fmt.Println("Контекст отменен:", ctx.Err()) // Можно узнать причину отмены
			return
		default:
			fmt.Println("Работаю с контекстом")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// 4. runtime.Goexit() - остановка текущей горутины
// Низкоуровневая функция, которая немедленно завершает текущую горутину
// Важно: выполняет все отложенные вызовы (defer), но НЕ завершает программу
// Опасность: может обойти нормальные пути очистки
func goexitExit(wg *sync.WaitGroup) {
	defer fmt.Println("Defer выполнится перед Goexit") // defer ВСЕГДА выполняется
	defer wg.Done()

	// Вложенная горутина
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Вложенная горутина вызывает Goexit")
		runtime.Goexit() // Завершает только ЭТУ вложенную горутину
		// Код после Goexit НЕ выполняется
		fmt.Println("Эта строка не выполнится во вложенной горутине")
	}()

	time.Sleep(2 * time.Second)
	// Основная горутина продолжает работу, т.к. Goexit был вызван во вложенной
	fmt.Println("Эта строка выполнится - Goexit был во вложенной горутине")
}

// 5. Закрытие канала данных (Channel closure)
// Идиоматичный способ для worker'ов, читающих из канала
// Плюсы: чистый код, автоматическое завершение
// Минусы: работает только с for-range по каналу
func channelClosureExit(dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Цикл for range автоматически завершается при закрытии канала
	// Это самая чистая идиома для обработчиков каналов
	for data := range dataChan {
		fmt.Printf("Обрабатываю данные: %d\n", data)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Канал закрыт, завершаю работу")
}

// 6. Таймаут (Timeout)
// Горутина завершается по истечении заданного времени
// Плюсы: контроль времени выполнения
// Минусы: жесткая привязка ко времени
func timeoutExit(wg *sync.WaitGroup) {
	defer wg.Done()

	// time.After возвращает канал, в который будет отправлено значение через указанное время
	timeout := time.After(2 * time.Second)
	// Ticker для периодических действий
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop() // Важно освободить ресурсы тикера

	for {
		select {
		case <-timeout: // Сработал таймаут
			fmt.Println("Время вышло!")
			return
		case <-ticker.C: // Периодическое выполнение
			fmt.Println("Тик...")
		}
	}
}

// 7. Сигналы ОС (OS signals)
// Горутина реагирует на системные сигналы (Ctrl+C, SIGTERM и т.д.)
// Используется для graceful shutdown приложений
func signalExit(sigChan chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case sig := <-sigChan: // Блокируемся до получения сигнала
		fmt.Printf("Получен сигнал: %v\n", sig)
		return
	case <-time.After(3 * time.Second): // Запасной таймаут на случай отсутствия сигнала
		fmt.Println("Таймаут ожидания сигнала")
		return
	}
}

// 8. Паника с восстановлением (Panic/Recover)
// Аварийное завершение с контролируемой обработкой
// Плюсы: можно обработать критические ошибки
// Минусы: нарушает нормальный поток выполнения
// Важно: recover работает ТОЛЬКО в defer
func panicExit(wg *sync.WaitGroup) {
	// Анонимная функция для перехвата паники
	defer func() {
		if r := recover(); r != nil { // recover перехватывает панику
			fmt.Println("Восстановлено после паники:", r)
		}
		wg.Done() // Всегда помечаем завершение
	}()

	counter := 0
	for {
		counter++
		fmt.Println("Счетчик:", counter)
		time.Sleep(200 * time.Millisecond)

		if counter == 3 {
			panic("Искусственная паника!") // Вызываем панику
		}

		if counter == 5 { // Этот код не будет достигнут
			break
		}
	}
}

// 9. Контекст с таймаутом (Context with timeout)
// Комбинация контекста и таймаута - самый частый сценарий в production
func contextTimeoutExit(wg *sync.WaitGroup) {
	defer wg.Done()

	// Контекст с таймаутом 1.5 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel() // Важно вызывать cancel для освобождения ресурсов, даже если таймаут сработал

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст с таймаутом истек:", ctx.Err())
			return
		default:
			fmt.Println("Работаю с таймаутом...")
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup // WaitGroup для синхронизации горутин

	fmt.Println("=== Демонстрация всех методов остановки горутин ===")

	// 1. Выход по условию
	fmt.Println("1. Выход по условию:")
	wg.Add(1) // Увеличиваем счетчик перед запуском горутины
	go conditionExit(&wg)
	wg.Wait() // Ждем завершения горутины

	// 2. Через канал уведомления
	fmt.Println("\n2. Через канал уведомления:")
	stopChan := make(chan struct{}) // Канал без данных, только для сигналов
	wg.Add(1)
	go channelExit(stopChan, &wg)
	time.Sleep(1 * time.Second) // Даем горутине поработать
	close(stopChan)             // Закрываем канал - все читающие горутины получают сигнал
	wg.Wait()

	// 3. Через контекст с отменой
	fmt.Println("\n3. Через контекст с отменой:")
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go contextExit(ctx, &wg)
	time.Sleep(1 * time.Second) // Даем поработать
	cancel()                    // Отменяем контекст - сигнал всем горутинам, его использующим
	wg.Wait()

	// 4. runtime.Goexit()
	fmt.Println("\n4. runtime.Goexit():")
	wg.Add(1)
	go goexitExit(&wg)
	wg.Wait()

	// 5. Закрытие канала данных
	fmt.Println("\n5. Закрытие канала данных:")
	dataChan := make(chan int, 5) // Буферизированный канал
	wg.Add(1)
	go channelClosureExit(dataChan, &wg)

	// Отправляем данные в канал
	for i := 1; i <= 3; i++ {
		dataChan <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(dataChan) // Закрываем канал - горутина автоматически завершится
	wg.Wait()

	// 6. Таймаут
	fmt.Println("\n6. Таймаут:")
	wg.Add(1)
	go timeoutExit(&wg)
	wg.Wait()

	// 7. Сигналы ОС
	fmt.Println("\n7. Сигналы ОС:")
	sigChan := make(chan os.Signal, 1) // Буферизированный канал для сигналов
	// Настраиваем перехват сигналов (SIGINT = Ctrl+C, SIGTERM = graceful shutdown)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go signalExit(sigChan, &wg)

	// Эмулируем отправку сигнала SIGINT (как будто пользователь нажал Ctrl+C)
	time.Sleep(500 * time.Millisecond)
	sigChan <- syscall.SIGINT
	wg.Wait()

	// Останавливаем перехват сигналов (хорошая практика)
	signal.Stop(sigChan)

	// 8. Паника с восстановлением
	fmt.Println("\n8. Паника с восстановлением:")
	wg.Add(1)
	go panicExit(&wg)
	wg.Wait()

	// 9. Контекст с таймаутом
	fmt.Println("\n9. Контекст с таймаутом:")
	wg.Add(1)
	go contextTimeoutExit(&wg)
	wg.Wait()

	fmt.Println("\n=== Все методы продемонстрированы ===")
}
