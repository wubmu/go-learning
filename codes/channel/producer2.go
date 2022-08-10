package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

// 产品
type Product struct {
	No string // 产品编号
}

// 生产者
type Producer struct {
	No int
}

// 消费者
type Consumer struct {
	No int
}

var Box = make(chan Product, 10)

func (p Producer) work() {
	for {
		product := Product{No: uuid.NewV4().String()}
		Box <- product
		fmt.Printf("生产者-%d 生产了 产品-%s\n", p.No, product.No)

		// 休息一会
		time.Sleep(time.Millisecond * 200)
	}
}

// 消费者
func (c Consumer) buy() {
	for {
		// 从盒子中取出
		product := <-Box
		fmt.Printf("消费者-%d 消费了 产品-%s\n", c.No, product.No)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	for i := 1; i <= 2; i++ {
		p := Producer{No: i}
		go p.work()
	}

	for i := 1; i <= 5; i++ {
		c := Consumer{No: i}
		go c.buy()
	}

	// 用来阻塞main函数，不让其结束
	wait := make(chan struct{})
	<-wait
}
