package main

import "fmt"

func escape1() *int {
	var a int = 1
	return &a
}

// 栈空间不足
func escape2() {
	s := make([]int, 0, 10000)
	for index, _ := range s {
		s[index] = index
	}
}

// 变量大小不确定
// 编译期间无法确定slice的长度，这种情况为了保证内存的安全，
//编译器也会触发逃逸，在堆上进行分配内存。直接s := make([]int, 10)不会发生逃逸
func escape3() {
	number := 10
	s := make([]int, number) // 编译期间无法确定slice的长度
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

// 4. 动态类型
//动态类型就是编译期间不确定参数的类型、参数的长度也不确定的情况下就会发生逃逸
// 空接口 interface{} 可以表示任意的类型，如果函数参数为 interface{}
//，编译期间很难确定其参数的具体类型，也会发生逃逸。
// fmt.Println(a ...interface{})函数参数为interface，编译器不确定参数的类型，会将变量分配到堆上
func escape4() {
	fmt.Println(1111)
}

// 闭包引用
func escape5() func() int {
	var i int = 1
	return func() int {
		i++
		return i
	}
}
func main() {
	escape1()
}
