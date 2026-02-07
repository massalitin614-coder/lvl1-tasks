package oldlogger

import "fmt"

type OldLogger struct {
	prefix string
}

// NewOldLogger - конструктор старого логгера
func NewOldLogger(prefix string) *OldLogger {
	return &OldLogger{prefix: prefix}
}

// Log - метод старого логгера
func (ol *OldLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", ol.prefix, message)
}
func (ol *OldLogger) LogWithDate(level string, message string) {
	// Очень старый формат вывода
	fmt.Printf("%s: [%s] %s\n", level, ol.prefix, message)
}

// Mетод, который возвращает количество логов (просто для примера)
func (ol *OldLogger) GetLogCount() int {
	return 42
}
