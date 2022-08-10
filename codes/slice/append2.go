package main

import "fmt"

func main() {
	ss := []int{1, 2, 3}
	for i := 4; i < 20; i++ {
		fmt.Printf("addr:%p,len:%d,cap:%d, pointer addr: %p \n", &ss, len(ss), cap(ss), ss)
		ss = append(ss, i)
	}
}
