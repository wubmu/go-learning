package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan struct{})

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- struct{}{}

			if i%2 == 1 {
				fmt.Println("go-1", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("g2", i)
			}
		}
	}()

	wg.Wait()
}
