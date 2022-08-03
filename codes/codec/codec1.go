package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{"zhangsan", 20}
	enc1 := gob.NewEncoder(os.Stdout)
	// 编码
	e := enc1.Encode(p1)
	fmt.Println(e)

	file1, _ := os.OpenFile("/tmp/person.gob", os.O_CREATE|os.O_WRONLY, 0644)
	enc2 := gob.NewEncoder(file1)
	enc2.Encode(p1)

}
