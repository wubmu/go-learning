package main

import "fmt"

func main() {
	// 声明切片
	var sli1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli1), cap(sli1), sli1)

	var sli2 = []int{} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli1), cap(sli2), sli2)

	var sli3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli3), cap(sli3), sli3)

	sli4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli4), cap(sli4), sli4)

}
