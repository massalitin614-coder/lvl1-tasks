package main

import (
	"fmt"
	"log"
)

// createHugeString создает строку указанного размера.
// Аргумент size определяет количество байт в результирующей строке.
func createHugeString(size int) string {
	// make([]byte, size) создает слайс байтов указанной длины.
	// Все элементы слайса инициализируются нулевыми значениями (0 для byte).
	return string(make([]byte, size))
}

// someFunc возвращает первые n символов из строки длиной 1024 байта.
func someFunc(n int) (string, error) {
	// 1 << 10 - это битовый сдвиг: 1 сдвигается на 10 битов влево.
	// 1 << 10 = 2^10 = 1024
	v := createHugeString(1 << 10)

	if len(v) < n {
		return "", fmt.Errorf("строка слишком короткая: %d < %d", len(v), n)
	}

	// Безопасно берем подстроку: первые n байт строки v.
	substring := v[:n]

	// РЕШЕНИЕ ПРОБЛЕМЫ УТЕЧКИ ПАМЯТИ:
	// Преобразуем срез строки в слайс байтов, а затем обратно в строку.

	justString := string([]byte(substring))

	// Возвращаем результат. Когда функция завершится:
	// 1. Локальная переменная v выйдет из области видимости
	// 2. Сборщик мусора удалит большую строку (1024 байта)
	// 3. В памяти останется только маленькая строка (n байтов)
	return justString, nil
}

func main() {

	n := 100
	justString, err := someFunc(n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Длина строки:", len(justString))
}
