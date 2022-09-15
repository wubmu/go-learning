package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 100)
	go func() {
		defer wg.Done()
		for i := 0; i < 200; i++ {

			if len(ch) < 100 {
				ch <- i
			} else {
				fmt.Println("溢出：", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		for {
			if data, ok := <-ch; ok {
				fmt.Print(data, " ")
			} else {
				break
			}
		}
	}()
	wg.Wait()
}
