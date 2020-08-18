package _2for_if_switch_goto

import (
	"fmt"
	"testing"
)

// for
func TestFor(t *testing.T) {

	for i := 0; i < 10; i++ {

	}

	for {
		fmt.Println()
	}

	s := "hello, 小萝卜"
	for i, v := range s {
		fmt.Printf("%d=%c \n", i, v)
	}

}

// if
func TestIf(t *testing.T) {
	s := true
	if s {
		fmt.Printf("")
	}

	if b := s; b {
		fmt.Printf("")
	}
}

// switch
func TestSwitch(t *testing.T) {

	// 类型断言
	switch returnType().(type) {
	case string:
		fmt.Println("it's a string")
	case int, float64:
		fmt.Println("it's a number")
	}

	// fallthrough 关键字 fallthrough 进行执行下一个case，
	//且fallthrough不会判断下一个case的条件。话句话说，不论下一个case是否被匹配，都会被执行
	s := 5
	switch s > 1 {
	case true:
		fmt.Println(">1")
		fallthrough
	case false:
		fmt.Println("<1")
	}

}

func returnType() interface{} {
	return "2"
}

// goto 跳至label位置, 相应的continue break也可以使用label
func TestGoto(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 'a'; j < 'z'; j++ {
			fmt.Printf("%d-%c \n", i, j)

			if j == 'd' {
				goto label
			}
		}
	}
label:
	fmt.Printf("jump to label")
}
