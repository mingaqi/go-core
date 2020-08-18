package _1basic

import (
	"testing"
)

// go中return不是原子操作, 是分为两步运行
// 第一步: 返回值赋值
// 第二步: 真正的真正的RET返回
// 函数中存在defer, defer执行的实际在第一步和第二步之间 并且是倒序执行
func TestDe(t *testing.T) {
	println(f1())
	println(f2())
	println(f3())
	println(f4())
}

// 5  defer执行返回值已经赋值完成
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值
	}()
	return x //返回值=x
}

// 6 x为返回值,
func f2() (x int) {
	defer func() {
		x++ // x是返回值, 修改的是返回值
	}()
	return 5 // 返回值=x
}

// 5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=y=x=5
}

// 5 函数传参是值传递, defer中x赋值运算是值copy的运算, 对原数据没有影响
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值=x=5
}
