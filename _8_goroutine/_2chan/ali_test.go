package _2chan

import (
	"fmt"
	"testing"
	"time"
)

// 模拟运维发布系统：目标10000物理服务器，更新应用（print hello world），发布策略最高并发20.
// 1. 协程
// 2. chan
// 3. 并发控制
func TestGo(t *testing.T) {
	// 机器数 协程数
	machine, gCount := 1000, 20
	c := make(chan int, gCount)

	// send
	go func() {
		for true {
			machine--
			if machine == 0 {
				break
			}
			c <- 1
		}
		close(c)
	}()

	// receive
	for i := 0; i < gCount; i++ {
		go deploy(c)
	}

	// waiting for all go
	time.Sleep(time.Second * 5)
}

func deploy(c <-chan int) {
	for {
		if _, ok := <-c; !ok {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println("hello world!")
	}

}
