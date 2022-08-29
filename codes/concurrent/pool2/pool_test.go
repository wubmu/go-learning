package pool2

import (
	"fmt"
	"sync"
	"testing"
)

func DemoFunc(x int32) {
	fmt.Println(x)
}
func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	p, _ := NewPool(5)

	runTimes := 10000

	for i := 0; i < runTimes; i++ {
		wg.Add(1)

		task := &Task{
			Handler: func(v ...interface{}) {
				wg.Done()
				fmt.Println(v)

			},
			Params: []interface{}{i, i * i, "hello"},
		}
		p.submit(task)
	}

	wg.Wait()

	// 安全关闭
	p.Close()

}
