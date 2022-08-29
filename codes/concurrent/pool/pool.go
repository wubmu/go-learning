package pool

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type sig struct{}
type f func() error

// Errors that are used throughout the Tunny API.
var (
	ErrPoolNotRunning = errors.New("the pool is not running")
	ErrJobNotFunc     = errors.New("generic worker not given a func()")
	ErrWorkerClosed   = errors.New("worker was closed")
	ErrJobTimedOut    = errors.New("job request timed out")
)

// Pool 协程池

type Pool struct {
	// 线程池的存储能力
	capacity int32

	// 正在运行的数量
	running int32

	// 设置worker的过期时间
	expiryDuration time.Duration

	// 存储可用的工人。
	workers []*Worker

	// 通知线程池关闭自己
	release chan sig

	lock sync.Mutex

	once sync.Once
}

func NewPool(size int) (*Pool, error) {
	return NewTimingPool(size, DefaultCleanIntervalTime)
}

func NewTimingPool(size, expiry int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if expiry <= 0 {
		return nil, ErrInvalidPoolExpiry
	}

	p := Pool{
		capacity: int32(size),
		//freeSignal:     make(chan sig, math.MaxInt32),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
		running:        0,
	}

	// 启动定时清理过期的worker
	//p.monitorAndClear()
	return &p, nil
}

// 提交任务到pool
func (p *Pool) Submit(task func()) error {

	if len(p.release) > 0 {
		return ErrPoolClosed
	}

	w := p.getWorker()
	w.task <- task
	return nil
}

func (p *Pool) getWorker() *Worker {
	var w *Worker
	// 标志变量，判断当前正在运行的worker数量是否已到达Pool的容量上限
	waiting := false
	// 加锁，检测队列中是否有可用worker，并进行相应操作
	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1

	fmt.Println("空闲worker数量:", n+1)
	fmt.Println("协程池现在运行的worker数量：", p.running)
	if n < 0 {
		// 判断运行worker数目已达到该Pool的容量上限，置等待标志
		waiting = p.Running() >= p.Cap()

		// 当前队列有可用的worker,从队列尾部取出一个
	} else {
		w = idleWorkers[n]
		idleWorkers[n] = nil
		p.workers = idleWorkers[:n]
	}

	// 检测完毕， 解锁
	p.lock.Unlock()

	if waiting {
		// 利用锁阻塞等待直到有空闲worker
		for {
			p.lock.Lock()
			idleWorkers = p.workers
			l := len(idleWorkers) - 1
			if l < 0 {
				p.lock.Unlock()
				continue
			}
			w = idleWorkers[l]
			idleWorkers[l] = nil
			p.workers = idleWorkers[:l]
			p.lock.Unlock()
			break
		}

		// 当前无空闲worker但是Pool还没有满，
		// 则可以直接新开一个worker执行任务
	} else if w == nil {
		w = &Worker{
			pool: p,
			task: make(chan func(), 1),
		}
		w.run()
		// 运行work加1
		p.addRunning(1)
	}
	return w
}

func (p *Pool) Release() {

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
