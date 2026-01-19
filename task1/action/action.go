package action

import (
	"fmt"
	"lvl1/task1/human"

	"time"
)

type Action struct {
	human.Human
	actionType string
	actionTime time.Time
	location   string
}

// Создаем действие
func NewAction(h *human.Human, actionType, location string) (*Action, error) {
	if h == nil {
		return nil, fmt.Errorf("human cannot be nil")
	}
	if err := validateAction(actionType, location); err != nil {
		return nil, fmt.Errorf("when creating an action: %w", err)
	}

	return &Action{
		Human:      *h,
		actionType: actionType,
		actionTime: time.Now(),
		location:   location,
	}, nil
}

func validateAction(actionType, location string) error {
	if actionType == "" {
		return fmt.Errorf("actionType is empty")
	}
	if location == "" {
		return fmt.Errorf("location is empty")
	}
	return nil
}

// Начинаем действие
func (a *Action) StartAction() string {
	a.actionTime = time.Now()
	return fmt.Sprintf("%s начинает действие %s в [%s]\n", a.Name(), a.actionType, a.location)
}

// Получаем информацию о действии
func (a *Action) GetActionInfo() string {
	return fmt.Sprintf("Действие: %s, Место: %s, Время: %s\n", a.actionType, a.location, a.actionTime.Format("15:04 02.01.2006"))
}

// Изменяем тип действия
func (a *Action) ChangeActionType(newActionType string) error {
	if newActionType == "" {
		return fmt.Errorf("new action type cannot de empty")
	}

	a.actionType = newActionType
	return nil
}

// Изменяем место действия
func (a *Action) ChangeLocation(newLocation string) error {
	if newLocation == "" {
		return fmt.Errorf("new location cannot be empty")
	}
	a.location = newLocation
	return nil
}

// Объедененный метод
func (a *Action) IntroduceWithAction() {
	a.Introduce()
	fmt.Printf("Сейчас выполняю действие: %s в %s\n", a.actionType, a.location)
	fmt.Printf("Время начала: %s\n", a.actionTime.Format("15:04"))
}

// Полная информация
func (a *Action) GetFullInfo() string {
	return fmt.Sprintf("%s(%d лет, %s) - выполняет %s в %s", a.Name(), a.Age(), a.Profession(), a.actionType, a.location)
}
