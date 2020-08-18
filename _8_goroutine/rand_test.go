package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 随机数需要设置种子, 否则每次随机数一样
func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(10))
	}
}
