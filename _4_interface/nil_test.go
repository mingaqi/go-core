package _4_interface

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int
	tel  string
}

//  eface:interface在golang内部其实是一个结构体，有_type,data两个指针组成
// https://studygolang.com/articles/1749
//
func TestNil(t *testing.T) {
	var man *Person = nil
	var ai interface{} = man
	var ei interface{} = nil
	fmt.Printf("ai == nil: %v\n", ai == nil)
	fmt.Printf("ai == ei: %v\n", ai == ei)
	fmt.Printf("ei == man: %v\n", ei == man)
	fmt.Printf("ei == nil: %v\n", ei == nil)
}

// https://studygolang.com/articles/1908
// 无论该指针的值是什么：(*interface{}, nil)，这样的接口值总是非nil的，即使在该指针的内部为nil。
func TestNil1(t *testing.T) {
	var val interface{} = (*interface{})(nil)
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}
}
