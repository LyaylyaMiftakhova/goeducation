package _return

import (
	"strconv"
	"strings"
)

func GetReversedInt(number int) int {
	strInput := strconv.Itoa(number)
	strLen := len(strInput)
	chars := make([]string, strLen)

	for i := 0; i < strLen; i++ {
		chars[i] = string(strInput[strLen-1-i])
	}

	reversed, _ := strconv.Atoi(strings.Join(chars, ""))
	return reversed
}
