package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)

	// 单生产者
	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 消费者
	go func() {
		for elem := range ch {
			fmt.Println(elem)
		}
	}()

	time.Sleep(2 * time.Second)
}
