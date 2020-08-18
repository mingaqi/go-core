package _3_struct

import (
	"fmt"
	"testing"
)

// 模拟继承  继承父类字段和方法
type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Printf("%s, 会动! \n", a.name)
}

type dog struct {
	feet int
	animal
}

func TestDog(t *testing.T) {
	dog := dog{4, animal{"来福"}}
	dog.move()
}
