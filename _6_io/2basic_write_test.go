package _6_io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

// 必须指定下面三个
// os.O_RDWR	读写
// os.O_WRONLY	只写
// os.O_RDONLY	只读
//
// os.O_APPEND 	追加
// os.O_CREATE	创建文件
// os.O_TRUNC	清空
func TestOpenFile(t *testing.T) {
	file, err := os.OpenFile("xx.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("occurred error:", err)
		return
	}
	defer file.Close()
	wr := bufio.NewWriter(file)
	wr.Write([]byte("嘿嘿\n"))
	wr.WriteString("呵呵\n")
	wr.Flush()
}

// copy
func TestCopy(t *testing.T) {
	// 只读打开源文件
	src, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("occurred error:", err)
		return
	}
	defer src.Close()

	// 写+创建 打开目标文件
	target, err := os.OpenFile("./copy.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("occurred error:", err)
		return
	}
	defer target.Close()

	// 执行copy
	io.Copy(target, src)
}
