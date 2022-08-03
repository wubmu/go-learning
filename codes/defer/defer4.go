//但是，大多数内置函数（除了copy和recover）的调用的返回结果都不可以舍弃
//（至少对于标准编译器1.18来说是如此）。 另一方面，我们已经了解到延迟函数调用的所有返回结果必须都舍弃掉。 所以，很多内置函数是不能被延迟调用的。
package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d"}
	defer fmt.Println(s) // [a x y d]
	// defer append(s[:1], "x", "y") // 编译错误
	defer func() {
		_ = append(s[:1], "x", "y")
	}()
}
