package deleteNumber

import (
	"errors"
)

func RemoveDigit(number, shouldDelete int) (int, error) {
	// Преобразование числа в слайс цифр
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}

	// Поиск и удаление цифры
	found := false
	result := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == shouldDelete {
			found = true
			continue
		}
		result = result*10 + digits[i]
	}

	// Если цифра для удаления не найдена, возвращаем ошибку
	if !found {
		return 0, errors.New("цифра для удаления не найдена")
	}

	// Инверсия результата и возвращение
	return reverseNumber(result), nil
}

func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}
