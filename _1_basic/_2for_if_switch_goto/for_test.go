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
		break
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

// range 陷阱
// 临时变量val每次的地址都不会变
// 所以 循环最后一次指向3的地址  导致map中所有元素都为3
func TestRange(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		//a := val
		m[key] = &val
		fmt.Printf("%p \n", &val)
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
