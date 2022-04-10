/**
1.数组不可动态变化问题，一旦声明了，其长度就是固定的。

2.数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销。

3.数组赋值问题，同样类型的数组（长度一样且每个元素类型也一样）才可以相互赋值，反之不可以。
*/
package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr)
	fmt.Println(arr)

	modifyArr2(&arr)
	fmt.Println(arr)

	//运行会报错：cannot use arr (type [5]int) as type [6]int in assignment
	//var arr =  [5] int {1, 2, 3, 4, 5}
	//var arr_1 [5] int = arr
	//var arr_2 [6] int = arr
}

func modifyArr(a [5]int) {
	a[1] = 20
}
func modifyArr2(a *[5]int) {
	a[1] = 21
}
