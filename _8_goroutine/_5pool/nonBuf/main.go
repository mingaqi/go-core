package main

import (
	"fmt"
	"time"
)

// ---------------------------test-----------------
func main() {

	//创建一个Task
	t := NewTa(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务  如果指定task个数会发生死锁
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	//启动协程池p
	p.Run()
}
