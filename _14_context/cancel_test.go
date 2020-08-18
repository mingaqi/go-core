package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 使用context控制子goroutine结束
func TestCancel(t *testing.T) {
	c, cancel := context.WithCancel(context.Background())
	go f(c, "rocky")
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("延迟一秒关闭如无name打印则证明子goroutine关闭成功")
}

func f(c context.Context, name string) {
	go f1(c, "lucy")
LOOP:
	for true {
		fmt.Println(name)
		select {
		case <-c.Done():
			break LOOP
		default:
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func f1(c context.Context, name string) {
LOOP:
	for true {
		fmt.Println(name)
		select {
		case <-c.Done():
			break LOOP
		default:
		}
		time.Sleep(500 * time.Millisecond)
	}
}
