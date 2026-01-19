package main

import (
	"fmt"
	"sync"
)

/*
//Более простой вариант
func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			fmt.Printf("%d\n", square)

		}(num)
	}

	wg.Wait()
}
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	//Создаем буферизированный канал,так как размер слайса небольшой
	squeresChan := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			squeresChan <- square
		}(num)
	}
	//Обязательно закрываем канал, и ждем завершение всех горутин
	go func() {
		wg.Wait()
		close(squeresChan)
	}()
	//При закрытии канала, происходит автоматический выход из цикла
	for result := range squeresChan {
		fmt.Println(result)
	}
}
