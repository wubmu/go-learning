package main

//一个被延迟调用的函数值可能是一个nil函数值。这种情形将导致一个恐慌

import "fmt"

func main() {
	defer fmt.Println("此行可以被执行到")
	var f func() // f == nil
	defer f()    // 将产生一个恐慌
	fmt.Println("此行可以被执行到2")
	f = func() {} // 此行不会阻止恐慌产生
}
