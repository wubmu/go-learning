package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	fmt.Println(v, &v[0])
	for i := range v {
		fmt.Println(v, &v[0])
		v = append(v, i)
	}
	fmt.Println(v, &v[0])
}
