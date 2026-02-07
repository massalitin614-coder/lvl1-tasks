package main

import (
	"fmt"
	"task21/adapter"
	"task21/newapp"
)

func main() {
	fmt.Println("=== Демонстрация паттерна Адаптер ===")
	fmt.Println()

	// 1. Создаем адаптер для старого логгера
	fmt.Println("1. Создаем адаптер для старого логгера:")
	loggerAdapter := adapter.NewOldLoggerAdapter("MyApp")
	fmt.Printf("   Адаптер создан: %T\n", loggerAdapter)
	fmt.Println()

	// 2. Создаем современное приложение с адаптером
	fmt.Println("2. Создаем современное приложение:")
	app := newapp.NewApp("ProductionApp", loggerAdapter)
	fmt.Printf("   Приложение создано: %v\n", app)
	fmt.Println()

	// 3. Используем приложение с адаптированным логгером
	fmt.Println("3. Запускаем обработку данных:")
	fmt.Println("   Вызываем app.ProcessData(\"test data\")")
	fmt.Println("   --- Логи ниже ---")

	err := app.ProcessData("test data")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println("   --- Конец логов ---")
	fmt.Println()

	// 4. Демонстрация WithFields
	fmt.Println("4. Демонстрация контекстного логирования:")
	contextLogger := loggerAdapter.WithFields(map[string]interface{}{
		"user_id":    12345,
		"request_id": "abc-123",
	})

	// Создаем приложение с контекстным логгером
	contextApp := newapp.NewApp("ContextApp", contextLogger)
	contextApp.ProcessData("context data")

	fmt.Println()
	fmt.Println("=== Анализ работы адаптера ===")
	fmt.Println()

	// 5. Проверяем, что адаптер реализует нужный интерфейс
	fmt.Println("5. Проверка реализации интерфейса:")
	var loggerInterface newapp.Logger
	loggerInterface = loggerAdapter // Это работает!
	fmt.Printf("   Адаптер реализует newapp.Logger: %v\n",
		loggerInterface != nil)
	fmt.Println()

	// 6. Сравнение старого и нового интерфейсов
	fmt.Println("6. Сравнение интерфейсов:")
	fmt.Println("   Старый интерфейс (OldLogger):")
	fmt.Println("     - Log(message string)")
	fmt.Println("     - LogWithDate(level, message string)")
	fmt.Println("     - Не возвращает ошибки")
	fmt.Println("     - Нет структурированного логирования")
	fmt.Println()
	fmt.Println("   Новый интерфейс (Logger):")
	fmt.Println("     - Info(message string, fields ...interface{}) error")
	fmt.Println("     - Error(err error, message string, fields ...interface{}) error")
	fmt.Println("     - Debug(message string, fields ...interface{}) error")
	fmt.Println("     - WithFields(fields map[string]interface{}) Logger")
	fmt.Println("     - Возвращает ошибки")
	fmt.Println("     - Поддерживает структурированное логирование")
	fmt.Println()

	// 7. Преобразования, которые делает адаптер
	fmt.Println("7. Преобразования адаптера:")
	fmt.Println("   Info() → форматирование → Log()")
	fmt.Println("   Error() → форматирование → LogWithDate()")
	fmt.Println("   WithFields() → сохранение полей в адаптере")
	fmt.Println("   Возврат ошибок → всегда nil (старый логгер не умеет в ошибки)")
}
