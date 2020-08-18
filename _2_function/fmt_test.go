package _2_function

import (
	"fmt"
	"strings"
	"testing"
)

//格式化字符串
/*
	%s	字符串
	%T 	类型
	%v	输出值而不是地址
	%+v 输出结构体添加字段名
	%#v 更加详细的信息
	%t  bool
	%q  该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	%   用来转义 %%->%


	%b		表示为二进制 无小数部分 二进制指数的科学计数法
	%e		科学计数法
	%c    	该值对应的unicode码值 rune单个字符  类似Java中char类型
	%d    	表示为十进制
	%o    	表示为八进制		定义 0777 八进制
	%x    	表示为十六进制，使用a-f		定义 0xff
	%X    	表示为十六进制，使用A-F
	%U    	表示为Unicode格式：U+1234，等价于"U+%04X"
	%p		指针
	https://www.cnblogs.com/yinzhengjie/p/7680829.html
*/
func TestFmt(t *testing.T) {

	println(strings.Count("abcdfefa", "f"))

	// 类型别名
	type myString string
	var i myString = "aaa"
	fmt.Printf("%T\n", i)
}
