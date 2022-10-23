package main

import (
	"fmt"
	"testing"
)

// 格式化输出
type Person struct {
	Name string
	Id   int
}

var person = &Person{Name: "Davie",
	Id: 1}

func TestPrintFmt(t *testing.T) {

	// 结构体
	fmt.Printf("%v \n", person)  // &{Davie 1}
	fmt.Printf("%+v \n", person) // &{Name:Davie Id:1}
	fmt.Printf("%#v \n", person) // &main.Person{Name:"Davie", Id:1}
	fmt.Printf("%T \n", person)  // *main.Person
	fmt.Printf("%% \n")          // %

	// go的原始类型
	fmt.Printf("%v \n", 123)  // 123
	fmt.Printf("%+v \n", 123) // 123
	fmt.Printf("%#v \n", 123) // 123
	fmt.Printf("%T \n", 123)  // int

	// go string
	fmt.Printf("%v \n", "123")  // 123
	fmt.Printf("%+v \n", "123") // 123
	fmt.Printf("%#v \n", "123") // "123"
	fmt.Printf("%T \n", "123")  // string
}

func TestPrintFmtInt(t *testing.T) {
	n := 97
	fmt.Printf("%b\n", n) // 1100001
	fmt.Printf("%c\n", n) // a
	fmt.Printf("%d\n", n) // 97
	fmt.Printf("%o\n", n) // 141
	fmt.Printf("%x\n", n) // 61
	fmt.Printf("%X\n", n) // 61
	fmt.Printf("%U\n", n) // u+0061
	fmt.Printf("%q\n", n) // 'a'

	m := 180
	fmt.Printf("%U\n", m)      // U+00B4
	fmt.Printf("%x\n", m)      // b4
	fmt.Printf("%X\n", m)      // B4
	fmt.Printf("%q\n", 0x0061) // 'a'
}

func TestPrintfFloat(t *testing.T) {
	f := 12.345

	fmt.Printf("%b\n", f)   // 6949617174986097p-49
	fmt.Printf("%e\n", f)   // 1.234500e+01
	fmt.Printf("%E\n", f)   // 1.234500E+01
	fmt.Printf("%f\n", f)   // 12.345000
	fmt.Printf("%.2f\n", f) // 12.35
	fmt.Printf("%F\n", f)   // 12.345000
	fmt.Printf("%g\n", f)   // 12.345
	fmt.Printf("%G\n", f)   // 12.345
}

func TestPrintString(t *testing.T) {
	s := "字符串"
	b := []byte{65, 66, 67}
	fmt.Printf("%s\n", s)   // 字符串
	fmt.Printf("%10s\n", s) //        字符串
	fmt.Printf("%s\n", b)   // ABC
	fmt.Printf("%b\n", b)   // [1000001 1000010 1000011]
	fmt.Printf("%v\n", b)   // [65 66 67]

	fmt.Printf("%q\n", s) //"字符串"
	fmt.Printf("%x\n", s) // e5ad97e7aca6e4b8b2
	fmt.Printf("%X\n", s) // E5AD97E7ACA6E4B8B2
}

func TestFmtPrintWeigh(t *testing.T) {
	m := 12.345
	fmt.Printf("%f\n", m)    // 12.345000
	fmt.Printf("%9f\n", m)   // 12.345000
	fmt.Printf("%9.2f\n", m) //     12.35
	fmt.Printf("%.2f\n", m)  // 12.35
	fmt.Printf("%9.f\n", m)  //        12
}

func TestFmtPrintFlag(t *testing.T) {
	n := 10
	m := -10
	fmt.Printf("% d,% d\n", n, m)

	ss := []string{"67", "68", "69"}
	fmt.Printf("% x\n", ss)
	fmt.Printf("%x\n", "67")
}
