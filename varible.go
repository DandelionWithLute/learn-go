package main

import (
	"fmt"
)

func varible() {
	var name1 string = "John"
	fmt.Println(name1)
	name2 := "Jane"
	fmt.Println(name2)

	age1 := 20
	fmt.Println(age1)
	age2 := 25
	fmt.Println(age2)

	mike, alice := 21, 22
	fmt.Println(mike, alice)
	fmt.Print(mike + alice)

	fmt.Printf("student %v is at the age of %v %T\n", name1, age1, age1)
	fmt.Printf("student %v is at the age of %v %T\n", name2, age2, age2)

	floatNum := 3.14
	fmt.Printf("%v %T", floatNum, floatNum)
}
