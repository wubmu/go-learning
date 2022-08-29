package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sort"
	"strconv"
	"testing"
)

func TestIterator(t *testing.T) {
	m := cmap.New[Animal]()

	// Insert 100 elements.
	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), Animal{strconv.Itoa(i)})
	}

	counter := 0
	// Iterate over elements.
	// m.Iter()
	for item := range m.IterBuffered() {
		val := item.Val
		key := item.Key
		fmt.Println(val, key)
		if (val == Animal{}) {
			t.Error("Expecting an object.")
		}
		counter++
	}

	if counter != 100 {
		t.Error("We should have counted 100 elements.")
	}
}

func TestConcurrent(t *testing.T) {
	m := cmap.New[int]()

	ch := make(chan int)
	const iterations = 1000
	var a [iterations]int

	// 适用goroutines 插入1000条数据
	go func() {
		for i := 0; i < iterations/2; i++ {
			m.Set(strconv.Itoa(i), i)

			// 再从map读出出来
			val, _ := m.Get(strconv.Itoa(i))
			ch <- val
		}
	}()

	go func() {
		for i := iterations / 2; i < iterations; i++ {
			// Add item to map.
			m.Set(strconv.Itoa(i), i)

			// Retrieve item from map.
			val, _ := m.Get(strconv.Itoa(i))

			// Write to channel inserted value.
			ch <- val
		} // Call go routine with current index.
	}()

	// 等待所有协程完成
	counter := 0
	for elem := range ch {
		a[counter] = elem
		counter++
		if counter == iterations {
			break
		}
	}
	// 对数组进行排序，将使验证我们返回的所有插入值变得更简单。
	sort.Ints(a[0:iterations])
}
