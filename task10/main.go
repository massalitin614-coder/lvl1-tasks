package main

import (
	"fmt"
	"slices"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	//Группируем температуру
	for _, temp := range temperatures {
		key := int(temp/10) * 10
		groups[key] = append(groups[key], temp)
	}

	//Сортируем по ключам
	keys := make([]int, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	fmt.Println("Выводим результат")
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}

}
