package main

import (
	"fmt"
	"sync"
	"time"
)

func channelClosureExit(wg *sync.WaitGroup) {
	defer wg.Done()

	timeout := time.After(2 * time.Second)
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			fmt.Println("Время вышло")
			return
		case <-ticker.C:
			fmt.Println("Тик")
		}
	}

}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Таймаут")
	wg.Add(1)
	go channelClosureExit(&wg)
	wg.Wait()
}
