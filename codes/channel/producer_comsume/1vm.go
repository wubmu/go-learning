package main

import "fmt"

func main() {
	var ch = make(chan int)

	// 单生产者
	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 多消费者
	for i := 0; i < 100; i++ {
		go func() {
			for elem := range ch {
				fmt.Println(elem)
			}
		}()
	}

	select {}
}
