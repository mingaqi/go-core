package main

import (
	"fmt"
	"go-core/_5_package/rectangle"
)

// 可以定义多个init函数, 执行顺序按照导入顺序执行, 是在main函数执行之前，自动被调用
var l, w float64 = 3.56, 4

func init() {
	fmt.Println("main package initialized")
}

func main() {
	dia := rectangle.Diagonal(l, w)
	area := rectangle.Area(l, w)
	fmt.Printf("%.2f,%.2f \n", dia, area)
}
