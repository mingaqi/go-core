package _1goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	Run()

}
