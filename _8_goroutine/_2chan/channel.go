package _2chan

import (
	"fmt"
	"sync"
)

/*producer 协程会从 0 到 9 写入信道 chn1，然后关闭该信道。
run有一个无限的 for 循环，使用变量 ok 检查信道是否已经关闭。
如果 ok 等于 false，说明信道已经关闭，于是退出 for 循环。
如果 ok 等于 true，会打印出接收到的值和 ok 的值*/
func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func Run() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

// 单向通道
func OnlyRead(c <-chan int, wg *sync.WaitGroup) {
	for {
		i, ok := <-c
		if ok {
			fmt.Println(i)
			wg.Done()
		}
	}
}
