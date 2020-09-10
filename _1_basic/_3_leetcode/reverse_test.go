package _3_leetcode

import (
	"fmt"
	"testing"
)

// 为什么超范围后是个复数
func TestReverse(t *testing.T) {
	fmt.Println(reverse(1534236469))
	temp := int32(1534236469)

	fmt.Println((temp*10)/10, temp, (temp*10)/10 == temp)

}
func reverse(x int) int {
	var res int

	for x != 0 {
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		y := x % 10
		res = res*10 + y
		x = x / 10
	}
	return res

}
