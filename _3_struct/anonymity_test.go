package _3_struct

import (
	"fmt"
	"testing"
)

// 使用匿名字段, 可以直接使用变量访问子结构体的字段, 类似于java中的继承功能
// 为防止结构体嵌套时 多个结构体字段冲突,
type company struct {
	Name      string `json:"name" ini:"name"`
	address   `json:"address"`
	workPlace `json:"workPlace"`
}
type address struct {
	Province string `json:province`
	City     string `json:city`
}
type workPlace struct {
	Province string `json:province`
	City     string `json:city`
}

func TestAnonymity(t *testing.T) {
	c := &company{
		"domi",
		address{"陕西", "西安"},
		workPlace{"广州", "珠海"},
	}
	c.address.Province = "山西"
	c.address.City = "运城"
	fmt.Printf("%#v \n", c)
}
