package main

import "fmt"

// Удаляет элемент из слайса значений (int)
func removeValue(slice []int, idx int) []int {
	if idx < 0 || idx >= len(slice) {
		return slice
	}

	copy(slice[idx:], slice[idx+1:])
	return slice[:len(slice)-1]
}

func main() {
	values := []int{100, 200, 300, 400, 500}
	index := 1
	fmt.Println("До удаления:", values)

	values = removeValue(values, index)

	fmt.Println("После удаления:", values)
}
