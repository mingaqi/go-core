package _6select

import (
	"fmt"
	"testing"
)

// select 可用于从多个channel中读数据,
func TestSelect(t *testing.T) {
	c1 := make(chan int, 5)
	fmt.Printf("channel CAP:%d ,len:%d\n", cap(c1), len(c1))

	for i := 0; i < 10; i++ {
		select {
		case x := <-c1:
			fmt.Println(x)
		case c1 <- i:
		default:
			fmt.Println("default")
		}
	}
}
