package _4_rune_char

import (
	"fmt"
	"testing"
)

// ascii
// https://segmentfault.com/q/1010000000404709
// rune=int32类型
func TestChar(t *testing.T) {
	var c rune
	c = 'c'
	fmt.Println(c)
	fmt.Println(string(c))
	fmt.Println(fmt.Sprintf("%d", 'a'))
	fmt.Println(fmt.Sprintf("%c", 65))
}
