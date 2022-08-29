package pool

import "time"

type Worker struct {
	// pool who owns this worker.
	pool *Pool

	// task is a job should be done.
	task chan func()

	// recycleTime 将在将工作人员放回队列时更新
	recycleTime time.Time
}

func (w *Worker) run() {
	go func() {
		// 监听任务队列，一旦有任务立马执行
		for f := range w.task {
			if f == nil {
				w.pool.addRunning(-1)
				return
			}
			f()
			// worker回收复用
			w.pool.putWorker(w)
		}
	}()
}

func (p *Pool) putWorker(worker *Worker) {

	// 写入放入时间
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.lock.Unlock()
}
