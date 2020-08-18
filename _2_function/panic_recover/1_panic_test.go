package panic_recover

import (
	"fmt"
	"testing"
)

// panic recover用来出来错误
// golang中没有异常机制  panic可以在任何地方引发, recover只能在defer调用的函数中有效

// panic错误一般是block级别错误, 用于阻断程序运行
// recover 用于恢复程序

func TestPanic(t *testing.T) {
	fA()
	fB()
	fC()
}

func fA() {
	fmt.Println("A")
}

func fB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("发生不可恢复错误")
		}
	}()
	panic("连接异常!!!")
	fmt.Println("B")
}

func fC() {
	fmt.Println("C")
}
