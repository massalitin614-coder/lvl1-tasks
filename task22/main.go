package main

import (
	"fmt"
	"log"
	"math/big"
)

type BigNumberCalculator struct {
	a *big.Int
	b *big.Int
}

// Создаем калькулятор
func NewBigNumberCalculator() *BigNumberCalculator {
	return &BigNumberCalculator{
		a: new(big.Int),
		b: new(big.Int),
	}
}

// Устанавливаем значение для а и b из строк в десятичной системе
func (bc *BigNumberCalculator) SetBigStrings(aStr, bStr string) error {
	a, ok := bc.a.SetString(aStr, 10)
	if !ok {
		return fmt.Errorf("неверный формат первого числа")
	}
	b, ok := bc.b.SetString(bStr, 10)
	if !ok {
		return fmt.Errorf("неверный формат второго числа")
	}
	bc.a = a
	bc.b = b
	return nil
}

// Устанавливаем значение напрямую из big.Int
func (bc *BigNumberCalculator) SetBigNumbers(a, b *big.Int) {
	bc.a = a
	bc.b = b
}

// Возращаем значение
func (bc *BigNumberCalculator) GetNumbers() (*big.Int, *big.Int) {
	return bc.a, bc.b
}

// Сложение
func (bc *BigNumberCalculator) Add() *big.Int {
	return new(big.Int).Add(bc.a, bc.b)
}

// Вычитание
func (bc *BigNumberCalculator) Sub() *big.Int {
	return new(big.Int).Sub(bc.a, bc.b)
}

// Умножение
func (bc *BigNumberCalculator) Mult() *big.Int {
	return new(big.Int).Mul(bc.a, bc.b)
}

// Деление, с проверкой второго числа
func (bc *BigNumberCalculator) Div() (*big.Int, error) {
	if bc.b.Sign() == 0 {
		return nil, fmt.Errorf("делить на ноль нельзя")
	}
	return new(big.Int).Div(bc.a, bc.b), nil
}

// Печатает текущие числа
func (calc *BigNumberCalculator) PrintNumbers() {
	fmt.Printf("a = %v\n", calc.a)
	fmt.Printf("b = %v\n", calc.b)
}
func main() {
	calc := NewBigNumberCalculator()
	err := calc.SetBigStrings("5000000000", "300000")
	if err != nil {
		log.Fatal(err)
	}
	calc.PrintNumbers()
	fmt.Println("Результаты операций")
	fmt.Printf("Сложение: %v\n", calc.Add())
	fmt.Printf("Вычитание: %v\n", calc.Sub())
	fmt.Printf("Умножение: %v\n", calc.Mult())
	result, err := calc.Div()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Деление: %v\n", result)

	fmt.Println("=======================")

	a := big.NewInt(100000000)
	b := big.NewInt(34543424)
	calc.SetBigNumbers(a, b)
	fmt.Printf("а = %v, b = %v\n", a, b)
	fmt.Printf("Сложение: %v\n", calc.Add())

}
