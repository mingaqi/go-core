package closure

import (
	"fmt"
	"strings"
	"testing"
)

func TestCall(t *testing.T) {
	Call()

	fmt.Println("----------------------------")

	match := matchSuffix(".jpg")
	fmt.Println(match("2014.jpg"))
	fmt.Println(match("花狗"))
}

// 闭包判断文件后缀
func matchSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
