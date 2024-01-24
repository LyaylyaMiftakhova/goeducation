package main

import (
	"example.com/goeducation/culculations"
	"example.com/goeducation/deleteNumber"
	"example.com/goeducation/myFirstServer"
	_return "example.com/goeducation/return"
	"fmt"
)

func main() {
	max := culculations.GetMax(2459234234)
	fmt.Println(max)

	reversed := _return.GetReversedInt(12345678)
	fmt.Println(reversed)

	reversed, err := deleteNumber.RemoveDigit(34567, 7)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(reversed)

	myFirstServer.MyFirstServer()
}
