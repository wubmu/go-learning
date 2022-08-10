package main

import (
	"fmt"
	"time"
)

var (
	infos = make(chan int, 10)
)

func producer(index int) {
	infos <- index
	fmt.Printf("producer : %d, send: %d\n", index, index)
}

func consumer(index int) {
	fmt.Printf("Consumer : %d, Receive: %d\n", index, <-infos)
}

// 生产者
func main() {
	//var wg sync.WaitGroup
	for index := 0; index < 10; index++ {
		go producer(index)
	}

	for index := 0; index < 10; index++ {
		go consumer(index)
	}

	time.Sleep(40 * time.Second)
}
