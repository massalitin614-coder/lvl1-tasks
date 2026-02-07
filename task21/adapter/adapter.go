package adapter

import (
	"fmt"
	"task21/newapp"
	"task21/oldlogger"
)

// OldLoggerAdapter - адаптер для старого логгера
// Реализует интерфейс newapp.Logger
type OldLoggerAdapter struct {
	// Встраиваем старый логгер
	*oldlogger.OldLogger
	fields map[string]interface{}
}

// Конструктор адаптера
func NewOldLoggerAdapter(prefix string) *OldLoggerAdapter {
	return &OldLoggerAdapter{
		OldLogger: oldlogger.NewOldLogger(prefix),
		fields:    make(map[string]interface{}),
	}
}

// Info - реализуем метод Info интерфейса Logger
func (a *OldLoggerAdapter) Info(message string, fields ...interface{}) error {
	// Преобразуем вызов Info в формат старого логгера

	// 1. Форматируем сообщение с полями
	formattedMsg := a.formatMessage("INFO", message, fields...)

	// 2. Вызываем старый метод Log
	// OldLogger.Log принимает только строку, не возвращает ошибку
	a.Log(formattedMsg)

	// 3. Возвращаем nil, так как старый логгер не возвращает ошибки
	return nil
}

// Error - реализуем метод Error интерфейса Logger
func (a *OldLoggerAdapter) Error(err error, message string, fields ...interface{}) error {
	// Форматируем сообщение с ошибкой
	formattedMsg := a.formatMessage("ERROR", fmt.Sprintf("%s: %v", message, err), fields...)

	// Используем старый метод LogWithDate
	a.LogWithDate("ERROR", formattedMsg)

	return nil
}

// Реализуем метод Debug интерфейса Logger
func (a *OldLoggerAdapter) Debug(message string, fields ...interface{}) error {
	formattedMsg := a.formatMessage("DEBUG", message, fields...)
	a.Log(formattedMsg)
	return nil
}

// Реализуем метод WithFields интерфейса Logger
func (a *OldLoggerAdapter) WithFields(fields map[string]interface{}) newapp.Logger {
	// Создаем новый адаптер с добавленными полями
	newAdapter := &OldLoggerAdapter{
		OldLogger: a.OldLogger, // Та же ссылка на старый логгер
		fields:    make(map[string]interface{}),
	}

	// Копируем существующие поля
	for k, v := range a.fields {
		newAdapter.fields[k] = v
	}

	// Добавляем новые поля
	for k, v := range fields {
		newAdapter.fields[k] = v
	}

	return newAdapter
}

// Вспомогательный метод для форматирования сообщения
func (a *OldLoggerAdapter) formatMessage(level, message string, fields ...interface{}) string {
	// Начинаем с уровня и сообщения
	result := fmt.Sprintf("%s: %s", level, message)

	// Добавляем поля из адаптера
	if len(a.fields) > 0 {
		result += " | "
		for k, v := range a.fields {
			result += fmt.Sprintf("%s=%v ", k, v)
		}
	}

	// Добавляем переданные поля
	if len(fields) > 0 {
		if len(a.fields) == 0 {
			result += " | "
		}
		// fields переданы как пары key, value
		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				key := fmt.Sprintf("%v", fields[i])
				value := fields[i+1]
				result += fmt.Sprintf("%s=%v ", key, value)
			}
		}
	}

	return result
}
