package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// go run main.go 生成文件
// go tool trace trace.out

// trace的使用
// http://wen.topgoer.com/docs/golangxiuyang/golangxiuyang-1cmeduvk27bo0#83b7ix
func main() {
	//创建trace文件
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}
