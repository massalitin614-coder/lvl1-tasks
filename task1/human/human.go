package human

import (
	"fmt"
)

// Human - родительская структура
type Human struct {
	name       string
	age        int
	profession string
}

// NewHuman - конструктор с валидацией
func NewHuman(name string, age int, profession string) (*Human, error) {
	if err := validateHuman(name, age, profession); err != nil {
		return nil, fmt.Errorf("when creating a person : %w", err)
	}

	return &Human{
		name:       name,
		age:        age,
		profession: profession,
	}, nil
}

// Валидация
func validateHuman(name string, age int, profession string) error {
	if name == "" {
		return fmt.Errorf("name is empty")
	}
	if age < 0 || age > 150 {
		return fmt.Errorf("incorrect age: %d", age)
	}
	if profession == "" {
		return fmt.Errorf("profession is empty")
	}
	return nil
}

// Методы Human
func (h *Human) Introduce() {
	fmt.Printf("Я %s, мне %d (лет,года,год) по профессии %s\n",
		h.name, h.age, h.profession)
}

func (h *Human) GrowAge() {
	h.age++
}

func (h *Human) SayHello() {
	fmt.Printf("Привет! Меня зовут %s\n", h.name)
}

// Геттеры
func (h *Human) Name() string {
	return h.name
}

func (h *Human) Age() int {
	return h.age
}

func (h *Human) Profession() string {
	return h.profession
}

func (h *Human) NewProfession(profession string) error {
	if profession == "" {
		return fmt.Errorf("profession cannot be empty")
	}
	h.profession = profession
	return nil
}

func (h *Human) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	h.name = name
	return nil
}

func (h *Human) SetAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("incorrect age: %d", age)
	}
	h.age = age
	return nil
}
