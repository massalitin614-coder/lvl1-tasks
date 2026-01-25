package main

import (
	"fmt"
)

// SetBit устанавливает i-й бит (нумерация с 0) в значение bitValue
// Возвращает ошибку при некорректных параметрах
func SetBit(n int64, i uint, bitValue int) (int64, error) {
	// Проверка корректности индекса бита
	if i >= 64 {
		return n, fmt.Errorf("индекс бита %d выходит за пределы int64 (максимум 63)", i)
	}

	// Проверка корректности значения бита
	if bitValue != 0 && bitValue != 1 {
		return n, fmt.Errorf("значение бита должно быть 0 или 1, получено %d", bitValue)
	}

	// Создаём маску
	mask := int64(1 << i)

	if bitValue == 1 {
		// Установка бита в 1
		return n | mask, nil
	} else {
		// Установка бита в 0
		return n &^ mask, nil
	}
}

func main() {
	// Пример из условия
	var num int64 = 5
	fmt.Printf("Исходное число: %d (%08b)\n", num, num)

	// Установка 1-го бита (индекс 1) в 0
	result, err := SetBit(num, 1, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("1-й бит → 0: %d (%08b)\n\n", result, result)

	// Тест с ошибками
	_, err = SetBit(num, 70, 1)
	fmt.Println("Тест с i=70:", err)

	_, err = SetBit(num, 2, 3)
	fmt.Println("Тест с bitValue=3:", err)
}
