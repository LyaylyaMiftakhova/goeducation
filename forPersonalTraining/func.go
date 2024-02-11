package forPersonalTraining

import (
	"errors"
	"fmt"
)

func Divide(x, y int) (res int, er error) {
	if y == 0 {
		return 0, errors.New("cannot divide on zero")
	}

	return x / y, nil
}

func name() {
	name := "John"
	age := 30

	// Используем fmt.Sprintf для форматирования строки
	message := fmt.Sprintf("Привет, меня зовут %s, и мне %d лет.", name, age)

	// Выводим отформатированную строку
	fmt.Println(message)
}
