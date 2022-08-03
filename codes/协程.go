package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func main() {
	// 并发执行程序
	go running()
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
	//time.Sleep(time.Second)
	wg.Wait()
}
