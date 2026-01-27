package main

import (
	"fmt"
	"strings"
)

func uniqueStrings(strings []string) []string {
	unique := []string{}
	setMap := make(map[string]struct{})

	for _, val := range strings {
		if _, ok := setMap[val]; !ok {
			setMap[val] = struct{}{}
			unique = append(unique, val)
		}
	}

	return unique

}

func uniqueString(sl1 []string) []string {
	unique := []string{}
	setMap := make(map[string]bool)

	for _, ele := range sl1 {
		if !setMap[ele] {
			setMap[ele] = true
			unique = append(unique, ele)
		}
	}

	return unique
}

func Unique[T comparable](items []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, item := range items {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func main() {
	//1.
	sl := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("Выводим полученный результат")

	result1 := uniqueStrings(sl)

	for i, val := range result1 {
		fmt.Printf("%d: %s\n", i+1, val)
	}

	//2.

	fmt.Println("Получаем уникальные элементы")

	result2 := uniqueString(sl)

	fmt.Printf("%s\n", strings.Join(result2, ", "))

	//3.

	fmt.Println("=================")

	result3 := Unique(sl)
	fmt.Println(result3)
}
