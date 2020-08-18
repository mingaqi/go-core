package _time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	now := time.Now()

	// 时间格式化, 后面的数值固定, 为golang出生时间
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	// 时间戳
	fmt.Println(now.Unix())           //s
	fmt.Println(now.UnixNano() / 1e6) //ms
	fmt.Println(now.UnixNano())       //ns

	// 时间计算
	fmt.Println("时间+", now.Add(24*time.Hour))
	fmt.Println("时间-", now.Add(-10*time.Minute))

	// 两个时间差值 内部使用sub()函数来实现
	time.Since(now).Milliseconds()

	// 时间戳转时间
	fmt.Println(time.Unix(1595836531, 0))

	// string 转 time
	temp, err := time.Parse("2006-01-02 15:04", "2020-07-27 16:01")
	fmt.Println(err)
	fmt.Println(temp)

	// sleep  直接使用n会报错, 类型不匹配
	n := 5
	time.Sleep(time.Duration(n) * time.Second)

	// 定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}
}
