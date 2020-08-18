package _2chan

import (
	"fmt"
	"sync"
	"testing"
)

// channel必须使用make初始化才能使用
var c chan int
var wg sync.WaitGroup

func TestChannel(t *testing.T) {
	c = make(chan int) // 不带缓冲区的channel
	wg.Add(1)

	go func() {
		for {
			fmt.Println(<-c)
			break
		}
		wg.Done()
	}()

	c <- 10 // 先写数据会发生阻塞 deadlock
	wg.Wait()
	close(c)

	/*--------------------------------------*/
	c = make(chan int, 10) // 带缓冲区channel
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)

	/*--------------------------------------*/
	// 单向channel 多用于函数参数中, 用于限制函数对channel的操作(只能只读和只写)
	var w sync.WaitGroup
	c = make(chan int, 10) // 带缓冲区channel
	for i := 0; i < 10; i++ {
		c <- i
		w.Add(1)
		if i == 9 {
			close(c)
		}
	}
	go OnlyRead(c, &w)
	w.Wait()

}

func TestRun(t *testing.T) {
	Run()
}
