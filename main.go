package main

import (
	"example.com/goeducation/culculations"
	"example.com/goeducation/deleteNumber"
	"example.com/goeducation/forPersonalTraining"
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

	//myFirstServer.MyFirstServer()

	result, err := forPersonalTraining.Divide(10, 0)
	fmt.Println(result, err)

	/*Реализация по ПП
	naturalIds := make([]int64, 48)
	adIds := make([]int64, 48)

	n := 48
	i := 3
	resultPP := promotion.InsertAds(n, i, [48]int64(naturalIds), adIds)

	fmt.Println(resultPP)
	*/

	//Задание по горутинам
	//goroutinesHomeWork.GoroutinesHomeWork()

	fmt.Println(forPersonalTraining.DomainForLocale("en", "ru"))
}
