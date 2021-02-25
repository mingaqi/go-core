package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// ./main -cpu=true -mem=true
// go tool pprof cpu.pprof
// top 10

// 	flat：当前函数占用CPU的耗时
// 	flat%：:当前函数占用CPU的耗时百分比
// 	sun%：函数占用CPU的耗时累计百分比
// 	cum：当前函数加上调用当前函数的函数占用CPU的总耗时
// 	cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比

// list logicCode 函数的详细分析

// gin中使用pprof http://liumurong.org/2019/12/gin_pprof/
func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}
