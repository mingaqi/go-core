package _1basic

import (
	"fmt"
	"testing"
)

// 使用_跳过某些值
const (
	a = iota // 0
	b        // 1
	_
	d        // 3
	e        // 4
	f = 100  // 100
	g = iota // 6
	h        // 7
	j = 200  // 200
	k        // 200
)

const (
	aa, bb  = iota + 1, iota + 2 // 1,2
	ccc, dd                      // 3,4
	ee, ff                       // 5,6
)

func TestIota(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(j)
	fmt.Println(k)

}
