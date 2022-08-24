package main

import (
	"fmt"
	"time"
)

// 多个生产者1个消费者
func main() {
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	const numSender = 1000
	// 生产者 sender
	for i := 0; i < numSender; i++ {
		go func(index int) {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- index:
				}
			}
		}(i)
	}

	// 接受者
	// the receiver
	go func() {
		for value := range dataCh {
			if value == 500 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}

			fmt.Println(value)
		}
	}()

	time.Sleep(time.Second * 10)
}
