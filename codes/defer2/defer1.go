package main

import "fmt"

/*
a: 2
a: 1
a: 0

b: 3
b: 3
b: 3
*/
func main() {
	func() {
		for i := 0; i < 3; i++ {
			//// 此i为形参i，非实参循环变量i。
			defer fmt.Println("a:", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			//i := i // 在下面的调用中，左i遮挡了右i。
			// <=> var i = i
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
}
