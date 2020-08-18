package closure

import (
	"fmt"
)

// 什么是闭包
// 闭包是一个函数, 这个函数包含了它外部作用域的一个变量

// 底层原理
// 1. 函数可以作为返回值
// 2. 函数查找变量的顺序是先查找内部, 然后查找外部

// 特性:
// 在accumulator变量的声明周期内, 闭包所应用的的外部变量value一直有效

func accumulate(value int) func() int {

	return func() int {
		value++
		return value
	}
}

func Call() {

	// 创建一个累加器, 初始值为1
	accumulator := accumulate(1)

	// 累加1并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	// 创建一个累加器, 初始值为1
	accumulator2 := accumulate(10)

	// 累加1并打印
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}
