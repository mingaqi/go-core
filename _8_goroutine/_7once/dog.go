package _7once

import (
	"fmt"
	"sync"
)

var once sync.Once
var d *Dog

type Dog struct {
	name, color string
}

func (d *Dog) Move() {
	fmt.Println(d.name, "moved!")
}

// once实现单例
func GetInstance() *Dog {
	fmt.Printf("%v\n", d)
	once.Do(func() {
		d = &Dog{"小黄", "白色"}
	})
	return d
}
