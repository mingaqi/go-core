package _3_struct

import (
	"fmt"
	"testing"
)

// 结构体占用一块练习的内存
// https://www.liwenzhou.com/posts/Go/10_struct/
func TestStruct(t *testing.T) {
	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "rocky"
	s.age = 20
	fmt.Printf("type: %T  value:%v\n", s, s)

	// 常用的方式 键值对初始化
	var p = &Person{name: "rocky", gender: "男"}
	fmt.Printf("%p\n", p)
	fmt.Printf("type:%T, value:%v \n", p, p)

	// p2是值 而不是指针 值列表初始化
	p2 := Person{"李磊", "男"}
	fmt.Printf("type:%T, value:%v \n", p2, p2)

	// new 返回指针 new初始化
	p1 := new(Person)
	p1.name = "lucy"
	p1.gender = "nv"
	fmt.Printf("type:%T, value:%#v\n", p1, p1)

	// 使用构造函数创建
	p3 := NewPerson("hury", "男")
	fmt.Printf("type:%T, value:%#v\n", p3, p3)
	p3.say()

}

type Person struct {
	name, gender string
}

// 构造函数, 约定俗成使用new开头
func NewPerson(name string, gender string) *Person {
	return &Person{name, gender}
}

// 方法 简单理解为java中成员方法 例如: getter setter, 由实例调用
// 接受者表示是调用该方法的具体类型变量 多用类型名称首字母小写
// 值调用者调用指针方法，也会改变你的结构内的变量值
func (p *Person) say() {
	fmt.Printf("%s 说, 我的性别是:%s \n", p.name, p.gender)
}
