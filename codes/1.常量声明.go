package main

import (
	"fmt"
)

func main() {
	//常量声明
	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name1, name2 string = "Tom", "Jay"
	fmt.Println(name1, name2)

	const name3, age1 = "Tom", 30
	fmt.Println(name3, age1)

}
