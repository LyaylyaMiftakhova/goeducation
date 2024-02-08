package promotion

import (
	"fmt"
)

func InsertAds(n int, i int, naturalIds [48]int64, adIds []int64) [48]int64 {
	// Проверка на корректность входных данных
	if i < 0 || i >= len(naturalIds) {
		fmt.Println("Некорректное значение i")
		return naturalIds
	}

	// Создание слайса для результата
	result := make([]int64, len(naturalIds))

	// Копирование "натуральной выдачи" до места вставки и вставка ПП-товара
	copy(result[:i], naturalIds[:i])
	copy(result[i:i+n], adIds)

	// Копирование "натуральной выдачи" после места вставки
	copy(result[i+n:], naturalIds[i+n:])

	return [48]int64(result)
}
