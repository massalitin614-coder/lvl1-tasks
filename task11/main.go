package main

import (
	"fmt"
)

func intersectionMap(sl1, sl2 []int) []int {

	set1 := make(map[int]bool)

	for _, val := range sl1 {
		set1[val] = true
	}

	resultSet := make(map[int]bool)
	result := []int{}
	for _, val := range sl2 {
		if set1[val] && !resultSet[val] {
			resultSet[val] = true
			result = append(result, val)
		}
	}

	return result

}

func main() {
	sl1 := []int{1, 2, 3}
	sl2 := []int{2, 3, 4}

	result := intersectionMap(sl1, sl2)

	fmt.Println(result)

}
