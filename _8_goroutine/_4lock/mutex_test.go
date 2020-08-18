package _4lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestRun(t *testing.T) {
	Run()
}

var x int

// sync.Mutex 互斥锁他能保证同时只有一个goroutine访问共享资源
func Run() {

	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	defer m.Unlock()
	wg.Done()
}
