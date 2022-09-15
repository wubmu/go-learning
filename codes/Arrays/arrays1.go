package main

func Array() {

	// 数组声明 ，
	// 长度不能留空
	// [...]自动判断长度
	var a [3]int = [3]int{
		1, 3, 4,
	}

	a[0] = 123

	var b = [...]int{
		1, 3, 4,
	}

	b[0] = 123

}
