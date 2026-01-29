package main

import (
	"fmt"
)

// quickSort - сортирует массив чисел по возрастанию
func quickSort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}

	// Выбираем элемент из середины как опорный (pivot)
	pivot := sl[len(sl)/2]

	left, right := 0, len(sl)-1
	for left <= right {
		for sl[left] < pivot {
			left++
		}
		for sl[right] > pivot {
			right--
		}
		// Если указатели не встретились - меняем элементы местами
		if left <= right {
			sl[left], sl[right] = sl[right], sl[left]
			left++
			right--
		}
	}
	if right > 0 {
		quickSort(sl[:right+1])
	}
	if left < len(sl)-1 {
		quickSort(sl[left:])
	}
	return sl
}

// binarySearch - ищет элемент в отсортированном массиве
func binarySearch(slice []int, target int) int {
	// Если массив пустой - элемент не найден
	if len(slice) == 0 {
		return -1
	}

	// Проверяем, что массив действительно отсортирован
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return -2 // Если не отсортирован, возвращаем ошибку
		}
	}

	// Устанавливаем границы поиска: левую и правую
	left, right := 0, len(slice)-1

	for left <= right {
		// Находим середину текущего диапазона
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {

	sl := []int{7, 2, 1, 5, 8, 6, 4, 57}
	// Сортируем массив
	result := quickSort(sl)
	fmt.Println(result)

	num := 4
	res := binarySearch(result, num)
	fmt.Println(res)
}
