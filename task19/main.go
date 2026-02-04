package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Возвращает перевернутую строку с поддержкой Unicode
func reverseString(str string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("empty string")
	}

	// Преобразуем в руны для работы с Unicode
	runes := []rune(str)
	result := make([]rune, len(runes))

	// Два указателя: i с начала, j с конца
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		result[i] = runes[j] // Копируем в обратном порядке
	}

	return string(result), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	result, err := reverseString(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // Вывод: главрыба
}
