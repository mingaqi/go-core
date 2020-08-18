package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
	"path"
	"time"
)

// 使用tail实现实时接收文件输入
func main() {
	p, _ := os.Executable()
	fmt.Println(os.Executable())
	fmt.Println(path.Dir(p))

	file := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // Whence之后，再偏移n个字符开始读，偏移量大于一行内容时换行继续计算, Whence:0跳到文件开始位置读 1跳到文件的当前位置? 2跳到文件最后，不读取文件里原有的内容，从新加入的开始读
		MustExist: false,                                // 文件是否必须存在
		Poll:      true,
	}

	tails, err := tail.TailFile(file, config)
	if err != nil {
		fmt.Println("openfile err:", err)
	}

loop:
	for {
		select {
		case s, ok := <-tails.Lines:
			if !ok {
				fmt.Printf("tail file close reopen, filename:%s", tails.Filename)
				continue loop
			}
			fmt.Printf("line:%s, time:%s\n", s.Text, s.Time)
		default:
			time.Sleep(time.Second * 1)
		}
	}
}
