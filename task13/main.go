package main

import "fmt"

func addSub(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func swapXOR(a, b int) (int, int) {
	a = a ^ b // a = a XOR b
	b = a ^ b // b = (a XOR b) XOR b = a
	a = a ^ b // a = (a XOR b) XOR a = b
	return a, b
}

func main() {
	//1. Сложение/вычетание
	a := 10
	b := 20
	x, y := addSub(a, b)
	fmt.Printf("1. Сложение/вычетание: a = %d, b = %d\n", x, y)
	//2. Замена, самый простой способ

	a1 := 15
	b1 := 25
	x1, y1 := swap(a1, b1)
	fmt.Printf("2. (a,b = b,a): a1 = %d, b1 = %d\n", x1, y1)

	//3. XOR обмен

	a2 := 35
	b2 := 40
	x2, y2 := swapXOR(a2, b2)

	fmt.Printf("3. XOR: a2 = %d, b2 = %d\n", x2, y2)

}
