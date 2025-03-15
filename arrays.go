package main

import (
	"fmt"
)

func arrays() {
	// fmt.Print("123")

	numArr := [5]int{1, 2, 3, 4, 5}
	fmt.Print(numArr)

	for i := 0; i < len(numArr); i++ {
		fmt.Print(numArr[i])
	}
}
