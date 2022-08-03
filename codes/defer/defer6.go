package main

// 1342

type T int

func (t T) M(n int) T {
	print(n)
	return t
}

func main() {
	var t T
	// t.M(1)是方法调用M(2)的属主实参，因此它
	// 将在M(2)调用被推入延迟调用队列时被估值。
	defer t.M(1).M(2)
	t.M(3).M(4)
}
