package _1basic

import (
	"fmt"
	"sort"
	"testing"
)

/*
append

切片可以理解为引用类型, 但又不是真正的引用类型, 它包含三个值, 切片第一个元素的底层数组地址, len cap
*/
func TestAppend(t *testing.T) {
	s := []string{"北京", "上海", "广州"}

	fmt.Printf("s:%v,len:%d, cap:%d \n", s, len(s), cap(s))

	/*
		 append追加元素时, 原底层数组放不下会新建数组并且已2倍扩容, 所以必须使用原变量来接收
		 切片底层数组容量小于1024 已2倍扩容
		 大于等于1024 则会循环增加原容量的1/4

		需要注意,切片扩容还会根据元素类型而做不同的处理,int和string处理方式就不同
	*/
	s = append(s, "西安")
	fmt.Printf("s:%v,len:%d, cap:%d \n", s, len(s), cap(s))

	s = append(s, "武汉", "成都")
	fmt.Printf("s:%v,len:%d, cap:%d \n", s, len(s), cap(s))

	ss := []string{"苏州", "杭州"}
	s = append(s, ss...) // ... 表示拆开
	fmt.Printf("s:%v,len:%d, cap:%d \n", s, len(s), cap(s))
}

/*
copy
*/
func TestCopy(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Printf("a:%v,len:%d, cap:%d \n", b, len(b), cap(b))
}

/*
del
切片操作是操作底层数组,  切片不存值, 只存第一个元素地址 len cap
*/
func TestDel(t *testing.T) {
	a := [...]int{1, 3, 5, 7, 5}
	s1 := a[:]

	s1 = append(s1[:1], s1[2:4]...)
	fmt.Printf("s1:%v,len:%d, cap:%d \n", s1, len(s1), cap(s1))
	fmt.Println("底层切片:", a)

	s1 = s1[:0]
	fmt.Printf("s1:%v,len:%d, cap:%d \n", s1, len(s1), cap(s1))
}

func TestSort(t *testing.T) {
	a := [...]int{1, 3, 5, 7, 9, 2, 3, 4}
	sort.Ints(a[:])
	fmt.Println(a)

	// :前代表位置 后代表值
	var c = [5]int{2: 3, 4: 67}
	fmt.Println(c)
}
