package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	done := make(chan struct{})
	go func() {
		time.Sleep(d)
		close(done)
	}()
	//блокурием горутину до получения стгнала
	<-done
}

func sleeps(d time.Duration) {
	<-time.After(d)
}

func sleepss(d time.Duration) {
	start := time.Now()
	for time.Since(start) < d {

	}
}

func main() {
	fmt.Println("Начало")
	sleep(5 * time.Second)
	fmt.Println("Конец. Прошло 5 секунд")

	fmt.Println("=============")
	fmt.Println("Начало")
	sleeps(3 * time.Second)
	fmt.Println("Конец, прошло 3 секунды")

	fmt.Println("=================")
	fmt.Println("Начало")
	sleepss(time.Second * 8)
	fmt.Println("Конец. Прошло 8 секунд")
}
