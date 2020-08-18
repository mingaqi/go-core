package _1string

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

/*
	字符串操作
*/
func TestString(t *testing.T) {

	// byte=int8 rune=int32(可表示中文)   切片byte len=9  rune len=3  rune类似于Java中的char
	b := "白萝卜"
	b1 := []rune(b)
	b1[0] = '红'
	fmt.Println(len(b1))
	fmt.Println(string(b1))

	s := fmt.Sprintf("%s|对|%s|说", "小明", "小红")
	fmt.Println(s)

	// strings 提供字符串操作方法 split join
	p := strings.Split(s, "|")
	fmt.Println(p)
	fmt.Println(strings.Join(p, "$"))

	// 统计string中中文字符个数
	b = "hello白萝卜"
	var i int
	for _, v := range b {
		if unicode.Is(unicode.Han, v) {
			i++
		}
	}
	fmt.Printf("中文字符[%d]个 \n", i)

}
