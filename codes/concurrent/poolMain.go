package main

import (
	"fmt"
	"hellogo/codes/concurrent/pool"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func demoFunc2() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("222222222222")
}
func main() {
	p, _ := pool.NewPool(5)

	runTimes := 300
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	demo2 := func() {
		demoFunc2()
		wg.Done()
	}
	go func() {
		for i := 0; i < runTimes; i++ {
			wg.Add(1)
			_ = p.Submit(syncCalculateSum)
		}
	}()

	go func() {
		for i := 0; i < runTimes; i++ {
			wg.Add(1)
			_ = p.Submit(demo2)
		}
	}()

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks.\n")
	// 等待所有任务执行完成
	wg.Wait()
}
