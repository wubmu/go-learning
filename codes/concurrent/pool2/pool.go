package pool2

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// errors
var (
	// return if pool size <= 0
	ErrInvalidPoolCap = errors.New("invalid pool cap")
	// put task but pool already closed
	ErrPoolAlreadyClosed = errors.New("pool already closed")
)

// 运行状态
const (
	RUNNING = iota
	STOP
)

// Task 任务
type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}

type Pool struct {
	capacity int32
	running  int32
	status   int32
	chTask   chan *Task

	lock sync.Mutex
}

func NewPool(capacity int32) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	p := &Pool{
		capacity: capacity,
		status:   RUNNING,
		chTask:   make(chan *Task, capacity),
	}
	return p, nil
}

func (p *Pool) submit(task *Task) error {
	p.lock.Lock()
	fmt.Println("协程池现在运行的worker数量：", p.running)
	defer p.lock.Unlock()
	if p.status == STOP {
		return ErrPoolAlreadyClosed
	}

	if p.Running() < p.Cap() {
		p.run()
	}
	if p.status == RUNNING {
		p.chTask <- task
	}

	return nil
}
func (p *Pool) run() {
	p.addRunning(1)
	go func() {
		defer func() {
			p.addRunning(-1)
			p.checkWorker() // // check worker avoid no worker running
		}()
		for {
			select {
			case task, ok := <-p.chTask:
				if !ok {
					return
				}
				task.Handler(task.Params...)

			}

		}

	}()

}

func (p *Pool) checkWorker() {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.Running() == 0 && len(p.chTask) > 0 {
		p.run()
	}
}

func (p *Pool) setStatus(status int32) bool {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.status == status {
		return false
	}

	p.status = status

	return true
}

// Close close pool graceful
func (p *Pool) Close() {

	if !p.setStatus(STOP) { // stop put task
		return
	}

	for len(p.chTask) > 0 { // wait all task be consumed
		time.Sleep(1e6) // reduce CPU load
	}

	close(p.chTask)
}
func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}
func (p *Pool) addRunning(delta int) {
	atomic.AddInt32(&p.running, int32(delta))
}
