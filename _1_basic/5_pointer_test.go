package _1basic

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
new & make:
https://www.liwenzhou.com/posts/Go/07_pointer/

unsafe: https://www.cnblogs.com/qcrao-2018/p/10964692.html
*/
func TestPointer(t *testing.T) {
	/*
		&: 取地址
		*: 根据地址取值  解指针
	*/
	n := 18
	p := &n
	fmt.Printf("%T \n", p)

	m := *p
	fmt.Printf("%T \n", m)

	a := new([5]int)
	(*a)[0] = 1
	fmt.Println(a)

	b := make([]int, 5, 5)
	fmt.Println(b)

	var ia int32 = 1
	var ib int16 = 1
	var ic rune = '你'
	fmt.Println(unsafe.Sizeof(ia))
	fmt.Println(unsafe.Sizeof(ib))
	fmt.Println(ic, unsafe.Sizeof(ic))
}
