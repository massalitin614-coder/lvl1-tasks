package newapp

import "fmt"

type Logger interface {
	// Info - логирование информационных сообщений
	Info(message string, fields ...interface{}) error

	// Error - логирование ошибок
	Error(err error, message string, fields ...interface{}) error

	// Debug - логирование отладочной информации
	Debug(message string, fields ...interface{}) error

	// WithFields - создание контекстного логгера
	WithFields(fields map[string]interface{}) Logger
}

// Современное приложение
type App struct {
	logger Logger
	name   string
}

// Конструктор приложения
func NewApp(name string, logger Logger) *App {
	return &App{
		name:   name,
		logger: logger,
	}
}

// Mетод приложения, который использует логгер
func (a *App) ProcessData(data string) error {
	a.logger.Info("Начало обработки данных",
		"app", a.name,
		"data_length", len(data))

	// Имитация обработки
	if data == "" {
		err := fmt.Errorf("пустые данные")
		a.logger.Error(err, "Ошибка обработки данных",
			"app", a.name)
		return err
	}

	// Логируем успешную обработку
	a.logger.Debug("Данные успешно обработаны",
		"app", a.name,
		"data", data[:min(10, len(data))])

	return nil
}

// min - вспомогательная функция
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
