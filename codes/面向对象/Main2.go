package main

import "fmt"
import "math"

type Point struct{ X, Y float64 }

//这是给struct Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {

	p := Point{1, 2}
	q := Point{4, 6}

	distanceFormP := p.Distance   // 方法值(相当于C语言的函数地址,函数指针)
	fmt.Println(distanceFormP(q)) // "5"
	fmt.Println(p.Distance(q))    // "5"

	//实际上distanceFormP 就绑定了 p接收器的方法Distance

	distanceFormQ := q.Distance   //
	fmt.Println(distanceFormQ(p)) // "5"
	fmt.Println(q.Distance(p))    // "5"

	//实际上distanceFormQ 就绑定了 q接收器的方法Distance
}
