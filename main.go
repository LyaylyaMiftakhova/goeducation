package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 56322547
	max := findMaxDigit(n)
	fmt.Println(max)
}

func findMaxDigit(n int) int {
	nStr := strconv.Itoa(n)

	var digitStrings []string

	for i := 0; i < len(nStr); i++ {
		digitStrings = append(digitStrings, string(nStr[i]))
	}

	var digits []int

	for _, digitStr := range digitStrings {
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			fmt.Printf("Ошибка при преобразовании строки %s в число: %v\n", digitStr, err)
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
