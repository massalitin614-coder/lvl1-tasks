package main

import "fmt"

func quickSort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}
	pivot := sl[len(sl)/2]
	left := []int{}
	right := []int{}
	equal := []int{}

	for _, val := range sl {
		switch {
		case val < pivot:
			left = append(left, val)
		case val > pivot:
			right = append(right, val)
		default:
			equal = append(equal, val)
		}
	}
	return append(append(quickSort(left), equal...), quickSort(right)...)

}

func main() {
	sl := []int{2, 4, 5, 65, 6674, 5653, 3, 45, 66, -4, 6}

	result := quickSort(sl)

	fmt.Printf("Длина: %d\n", len(result))

	for ind, ele := range result {
		fmt.Printf("%d: %d\n", ind, ele)
	}
}
