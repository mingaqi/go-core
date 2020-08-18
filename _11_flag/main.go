package main

import (
	"flag"
	"fmt"
	"time"
)

// -flag xxx （使用空格，一个-符号）
// --flag xxx （使用空格，两个-符号）
// -flag=xxx （使用等号，一个-符号）
// --flag=xxx （使用等号，两个-符号）
// 其中，布尔类型的参数必须使用等号的方式指定

// ./main --help  查看帮助

// ./main -name=rocky -age=31 -married=true -cTime=10h  ab cd ef
func main() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var cTime time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&cTime, "cTime", 10, "结婚时间")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, cTime)

	//返回命令行参数后的其他参数 [ab cd ef]
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数 3
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数      4
	fmt.Println(flag.NFlag())

}
