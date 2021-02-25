package _2chan

import (
	"fmt"
	"math/rand"
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

	c <- 10 // 先写数据会发生阻塞 deadlock  无缓存channel写入数据会发生阻塞
	wg.Wait()
	close(c)

	/*--------------------------------------*/
	c = make(chan int, 10) // 带缓冲区channel
	c <- 1
	c <- 2
	c <- 3
	close(c) // close后还能接受数据
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

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

//  一个协程生产随机数  一个协程读   借助channel实现
func TestRand(t *testing.T) {
	output := make(chan int)
	var wg sync.WaitGroup

	randFunc := func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			output <- rand.Intn(10)
		}
		close(output)
	}

	wg.Add(2)
	go randFunc()
	go func() {
		defer wg.Done()

		for i := range output {
			fmt.Println(i)
		}

	}()
	wg.Wait()
}
