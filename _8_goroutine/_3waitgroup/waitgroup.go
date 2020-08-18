package _3waitgroup

import (
	"fmt"
	"sync"
	"time"
)

// 类似于Java中的CountDownLatch
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func Run() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
