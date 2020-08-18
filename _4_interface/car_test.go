package _4_interface

import (
	"fmt"
	"testing"
)

// 接口实现多态
// 如果实现了接口中所有的方法, 那么是实现了这个接口
// 一个类型可以实现多个接口
// 接口也可以想struct一样进行嵌套

type car interface {
	run()
}

type toyoto struct {
	brand string
}

func (t *toyoto) run() {
	fmt.Printf("%s, 极速:150 \n", t.brand)
}

type fararri struct {
	brand string
}

func (f *fararri) run() {
	fmt.Printf("%s, 极速:350 \n", f.brand)
}

func driver(c car) {
	c.run()
}

func TestCar(t *testing.T) {
	f := &fararri{
		"法拉利",
	}
	to := &toyoto{
		"丰田",
	}
	driver(f)
	driver(to)
}
