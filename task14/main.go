package main

import (
	"fmt"
	"reflect"
)

// Определяем стандартные типы через switch
func typeDetector(i interface{}) string {
	fmt.Printf("Анализируем: %v\n", i)
	switch t := i.(type) {
	case int:
		return fmt.Sprintf("Тип int: %d", t)
	case float64:
		return fmt.Sprintf("Тип float64: %.1f", t)
	case string:
		return fmt.Sprintf("Тип string: %s, длина: %d", t, len(t))
	case bool:
		return fmt.Sprintf("Тип bool: %t", t)
	case []int:
		return fmt.Sprintf("Тип []int: %v", t)
	default:
		return fmt.Sprintf("Неизвестный тип: %v", t)

	}
}

// Для каналов лучше использовать reflect, так как switch не распознаает все каналы
func typeDetectorChannel(v interface{}) string {

	t := reflect.TypeOf(v)

	if t == nil {
		return "nil"
	}

	if t.Kind() == reflect.Chan {
		return fmt.Sprintf("Канал: %v (элементы: %v)", t, t.Elem())
	}

	return "не канал"
}

func main() {

	slice := []interface{}{
		"hello",
		23,
		4.5,
		true,
		[]int{1, 2, 3},
		[]string{},
	}

	for i, ele := range slice {
		fmt.Printf("%d: %v\n", i+1, typeDetector(ele))
	}
	fmt.Println("===================")
	sliceChannels := []interface{}{
		make(chan int),
		make(chan []string),
		make(chan struct{}),
		make(chan func()),
		make(chan bool),
		make(chan<- string),
	}

	for i, val := range sliceChannels {
		fmt.Printf("%d: %v\n", i+1, typeDetectorChannel(val))
	}

}
