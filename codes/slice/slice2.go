package main

import "fmt"

// 数组
func main() {
	a2 := [5]int{1, 2, 3}
	fmt.Println(cap(a2))
	a1 := [...]int{1, 2, 3}
	fmt.Println(cap(a1))
}
