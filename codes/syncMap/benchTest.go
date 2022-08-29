package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"strconv"
	"sync"
	"time"
)

func main() {
	count := 10000000
	loop := 5

	startT := time.Now()
	cmap := cmap.New()
	for i := 0; i < count; i++ {
		cmap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
	fmt.Printf("cmap 写 time cost = %v\n", time.Since(startT))

	startT = time.Now()
	var m sync.Map
	for i := 0; i < count; i++ {
		m.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	fmt.Printf("sync.map 写 time cost = %v\n", time.Since(startT))

	startT = time.Now()
	for j := 0; j < loop; j++ {
		for i := 0; i < count; i++ {
			cmap.Get(strconv.Itoa(i))
		}
	}
	fmt.Printf("cmap 读 time cost = %v\n", time.Since(startT))

	startT = time.Now()
	for j := 0; j < loop; j++ {
		for i := 0; i < count; i++ {
			m.Load(strconv.Itoa(i))
		}
	}
	fmt.Printf("sync.map 读 time cost = %v\n", time.Since(startT))

	startT = time.Now()
	for i := count; i < count*loop; i++ {
		cmap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
	fmt.Printf("cmap 写 time cost = %v\n", time.Since(startT))

	startT = time.Now()
	for i := count; i < count*loop; i++ {
		m.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	fmt.Printf("sync.map 写 time cost = %v\n", time.Since(startT))
}
