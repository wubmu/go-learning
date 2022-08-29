package main

import (
	"fmt"
	"sync"
	"time"
)

//worker工作的最大时长，超过这个时长worker自行收工无需等待manager叫停
const MAX_WORKING_DURATION = 5 * time.Second

//达到实际工作时长后，manager可以提前叫停
const ACTUAL_WORKING_DURATION = 10 * time.Second

type ctx struct {
	mu     sync.Mutex
	closed bool
	done   chan struct{}
}

func New() (*ctx, func()) {
	c := ctx{}
	c.done = make(chan struct{})
	return &c, func() { c.cancel() }
}

func NewWithTimeout(dur time.Duration) (*ctx, func()) {
	c := ctx{}
	c.done = make(chan struct{})
	timer := time.NewTimer(dur)
	go func() {
		select {
		case <-timer.C:
			c.cancel()
		}
	}()

	return &c, func() { c.cancel() }
}

func (c *ctx) Done() chan struct{} {
	return c.done
}

func (c *ctx) cancel() {
	c.mu.Lock()
	if c.closed == true {
		c.mu.Unlock()
		return
	}
	close(c.done)
	c.closed = true
	c.mu.Unlock()
}

func main() {
	c, cancelFunc := NewWithTimeout(MAX_WORKING_DURATION)

	go worker(c, "[1]")
	go worker(c, "[2]")

	go manager(c, cancelFunc)

	<-c.Done()

	//暂停1秒便于协程的打印输出
	time.Sleep(1 * time.Second)
	fmt.Println("company closed")
}

func manager(c *ctx, cancelFunc func()) {
	time.Sleep(ACTUAL_WORKING_DURATION)
	fmt.Println("manager called cancel()")
	cancelFunc()
}

func worker(c *ctx, name string) {
	for {
		select {
		case <-c.Done():
			fmt.Println(name, "return for ctxWithCancel.Done()")
			return
		default:
			fmt.Println(name, "working")
		}
		time.Sleep(1 * time.Second)
	}
}
