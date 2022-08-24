package main

import "fmt"

//用channel和两个goroutine实现数组相加
func add(a, b []int) []int {
	ch := make(chan int)
	c := make([]int, len(a))
	go func() {
		for _, v := range a {
			ch <- v
		}
	}()
	go func() {
		for i, t := range b {
			temp := <-ch
			c[i] = temp + t
		}
	}()
	return c
}

func main() {
	a := []int{2, 4, 6, 8}
	b := []int{1, 3, 5, 7}
	ans := add(a, b)
	fmt.Println(ans)
}
