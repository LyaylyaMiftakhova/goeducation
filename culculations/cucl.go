package culculations

import (
	"fmt"
	"strconv"
)

func GetMax(a int) int {
	aStr := strconv.Itoa(a)

	var digitStrings []string

	for i := 0; i < len(aStr); i++ {
		digitStrings = append(digitStrings, string(aStr[i]))
	}

	var digits []int

	for _, digitStr := range digitStrings {
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			fmt.Printf("Ошибка при преобразовании строки %s в число: %v\n", digitStr, err)
			continue
		}
		digits = append(digits, digit)
	}

	max := 0

	for _, digit := range digits {
		if digit > max {
			max = digit
		}
	}

	return max
}
