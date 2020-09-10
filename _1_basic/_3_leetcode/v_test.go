package _3_leetcode

import (
	"fmt"
	"testing"
)

func TestV(t *testing.T) {
	var k = 9
	for k = range []int{} {

	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {

	}
	fmt.Println(k)

	for k = range (*[3]int)(nil) {

	}
	fmt.Println(k)
}

func Test2(t *testing.T) {

	/*
		    list := new([]int)
			list = append(list, 1) // new返回指针 append接收切片
			fmt.Println(list)
	*/

	/*
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5}
		s1 = append(s1, s2) // 使用...
	*/

}

//当函数返回值有一个命名的，如sum int，则所有的返回值都必须命名。
/*func funcMui(x, y int) (sum int, error) {
	return x + y, nil
}*/

//结构体只能比较是否相等，不能比较大小
//相同类型的结构体才能比较
//如果struct的所有成员都是可以比较的就可以通过==或!=进行比较是否相等。（切片、map、函数不能比较）
func TestStruct(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 = = sn2")
	}

	/*sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 = = sm2")
	}*/

}

// 补码 https://studygolang.com/articles/24985?fr=sidebar
// https://blog.csdn.net/zl10086111/article/details/80907428
func TestBj(t *testing.T) {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y)
}
