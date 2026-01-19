package main

import (
	"fmt"
	"log"
	"lvl1/task1/action"
	"lvl1/task1/human"
)

func main() {
	//Создаем Human через конструктор
	human, err := human.NewHuman("Андрей", 33, "программист")
	if err != nil {
		log.Fatal(err)
	}
	human.SayHello()
	human.Introduce()
	human.GrowAge()
	fmt.Println("Поздравьте у меня сегодня день рождения!!!")
	fmt.Printf("Теперь мне %d\n", human.Age())
	fmt.Println()

	//Создаем Action через конструктор
	action, err := action.NewAction(human, "написание кода", "офис")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("1. Унаследованные методы от Human:")
	fmt.Printf("Имя: %s\n", action.Name())
	fmt.Printf("Возраст: %d\n", action.Age())
	fmt.Printf("Профессия: %s\n", action.Profession())
	action.Introduce()
	fmt.Println()

	fmt.Println("2. Методы Action:")
	action.StartAction()
	fmt.Println(action.GetActionInfo())
	fmt.Println()

	fmt.Println("3. Комбинированные методы:")
	action.IntroduceWithAction()
	fmt.Println()

	fmt.Println("4. Изменение данных Action:")
	err = action.ChangeActionType("тестирование кода")
	if err != nil {
		log.Fatal(err)
	}

	err = action.ChangeLocation("удаленно")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(action.GetFullInfo())
	fmt.Println()

	fmt.Println("5. Работа с методами Human через Action:")
	// Вызываем GrowAge (метод Human)
	action.GrowAge()
	fmt.Printf("Новый возраст после дня рождения: %d\n", action.Age())

	// Вызываем NewProfession (метод Human)
	err = action.NewProfession("доставщик")
	if err != nil {
		log.Fatal(err)
	}

	// Выводим обновленную информацию
	action.IntroduceWithAction()

}
