package main

import "fmt"

// 无缓存channel 可重用goroutine的协程池
// https://www.jianshu.com/p/508f5d3b2f59
//-----------------------------task------------------------------
type Task struct {
	f func() error //一个无参函数类型
}

// 通过 new() 来创建任务
func NewTa(f func() error) *Task {
	return &Task{f: f}
}

// 执行task
func (t *Task) execute() {
	t.f()
}

//-----------------------------pool------------------------------
// 定义池类型
type Pool struct {
	// task接收
	EntryChannel chan *Task

	worker_num int

	jobsChannel chan *Task
}

// pool 构造函数
func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		worker_num:   cap,
		jobsChannel:  make(chan *Task),
	}
}

// 从channel中取任务并且执行
func (p *Pool) worker(work_ID int) {
	for task := range p.jobsChannel {
		task.execute()
		fmt.Println("worker ID ", work_ID, " 执行完毕任务")
	}
}

// 启动协程池
func (p *Pool) Run() {
	// 启动固定数量的worker并执行
	for i := 0; i < p.worker_num; i++ {
		go p.worker(i)
	}
	for task := range p.EntryChannel {
		p.jobsChannel <- task
	}
	close(p.EntryChannel)
	close(p.jobsChannel)
}
