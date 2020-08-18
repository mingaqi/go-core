package _4lock

import (
	"fmt"
	"sync"
	"testing"
)

// 读写锁 应对读多写少的情况 , 使用sync.RwMutex
// 读锁: 当一个goroutine获取了读锁, 其他的goroutine如果获取读锁会继续获取锁, 如果获取写锁会进入等待
// 写锁: 当一个goroutine获取了写锁, 其他goroutine无论获取读锁或写锁都会等待
func TestRw(t *testing.T) {
	for i := 0; i < 100; i++ {
		sg.Add(1)
		go read()
	}
	for i := 0; i < 50; i++ {
		sg.Add(1)
		go write()
	}
	sg.Wait()
}

var (
	a  int
	rw sync.RWMutex
	sg sync.WaitGroup
)

func read() {
	rw.Lock()
	a++
	fmt.Println("write:", a)
	rw.Unlock()
	sg.Done()
}
func write() {
	rw.RLock()
	fmt.Println("read:", a)
	rw.RUnlock()
	sg.Done()
}
