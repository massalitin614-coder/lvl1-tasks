package main

import (
	"fmt"
)

func reverseWords(str string) string {
	strRune := []rune(str)
	n := len(strRune)
	//Переворачиваем всю строку
	reverse(strRune, 0, n-1)
	//Переворачиваем каждое слово обратно
	start := 0
	for i := 0; i <= n; i++ {
		if i == n || strRune[i] == ' ' {
			reverse(strRune, start, i-1)
			start = i + 1
		}

	}
	return string(strRune)

}

// reverse переворачивает часть слайса от left до right включительно
func reverse(strRune []rune, left, right int) {
	for left < right {
		strRune[left], strRune[right] = strRune[right], strRune[left]
		left++
		right--
	}

}

func main() {
	input := "snow dog sun"
	fmt.Printf("Вход:  %q\n", input)
	fmt.Printf("Выход: %q\n", reverseWords(input))
}
