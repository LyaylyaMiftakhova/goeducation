package deleteNumber

import (
	"errors"
)

func RemoveDigit(number, shouldDelete int) (int, error) {
	// Поиск и удаление цифры, а также преобразование числа в слайс цифр
	found := false
	result := 0
	for number > 0 {
		digit := number % 10
		number /= 10

		if digit == shouldDelete {
			found = true
			continue
		}

		result = result*10 + digit
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
