package main

import (
	"fmt"
)

func main() {
	var age1 uint8 = 31
	var age2 = 32
	age3 := 33
	fmt.Println(age1, age2, age3)

	var age4, age5, age6 int = 31, 32, 33
	fmt.Println(age4, age5, age6)

	var name1, age7 = "Tom", 30
	fmt.Println(name1, age7)

	name2, isBoy, height := "Jay", true, 180.66
	fmt.Println(name2, isBoy, height)
}
