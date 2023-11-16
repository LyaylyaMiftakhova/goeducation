package main

import (
	"example.com/goeducation/culculations"
	"example.com/goeducation/return"
	"fmt"
)

func main() {
	max := culculations.GetMax(2459234234)
	fmt.Println(max)

	reversed := _return.GetReversedInt(12345678)
	fmt.Println(reversed)
}
