package _9atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// atomic https://www.liwenzhou.com/posts/Go/14_concurrence/
// sync.once中也使用atomic来实现
func TestAtomic(t *testing.T) {
	var a int32
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(&a, &wg)

	}
	wg.Wait()
	fmt.Println(a)

}

func add(a *int32, wg *sync.WaitGroup) {
	atomic.AddInt32(a, 1)
	(*wg).Done()
}
