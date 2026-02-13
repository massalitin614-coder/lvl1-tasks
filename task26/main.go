package main

import (
	"fmt"
	"unicode"
)

// Разбиваем строку по руна, приводим к нижнему регистру
// Добавдяем ключ в map, если попадается повторно выводит false
func uniqueSymbol(str string) bool {
	if len(str) < 2 {
		return true
	}
	seenRunes := make(map[rune]struct{})

	for _, r := range str {
		lower := unicode.ToLower(r)
		if _, ok := seenRunes[lower]; ok {
			return false
		}
		seenRunes[lower] = struct{}{}
	}
	return true
}

func main() {

	fmt.Println(uniqueSymbol("abcd"))
	fmt.Println(uniqueSymbol("abCdefAaf"))
	fmt.Println(uniqueSymbol("aabcd"))
	fmt.Println(uniqueSymbol(""))
	fmt.Println(uniqueSymbol("Aa"))
}
