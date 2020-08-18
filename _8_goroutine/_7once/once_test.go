package _7once

import (
	"fmt"
	"testing"
)

// once 只被执行一次 例如: 配置文件加载 , 关闭通道 , 单例
func TestOnce(t *testing.T) {
	d1 := GetInstance()
	d2 := GetInstance()
	fmt.Printf("%p\n", d1)
	fmt.Printf("%p\n", d2)

}
