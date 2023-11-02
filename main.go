package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 56322547
	max := 0

	nStr := strconv.Itoa(n)

	var digitStrings []string

	for i := 0; i < len(nStr); i++ {
		digitStrings = append(digitStrings, string(nStr[i]))
	}

	var digits []int

	for _, digitStr := range digitStrings {
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
		}
		digits = append(digits, digit)
	}

	for _, digit := range digits {
		if digit > max {
			max = digit
		}
	}
	fmt.Println(max)
}
