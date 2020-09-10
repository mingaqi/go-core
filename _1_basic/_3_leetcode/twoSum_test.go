package _3_leetcode

import (
	"fmt"
	"testing"
)

// 两数之和 每个数不能使用第二次
func Test2Sun(t *testing.T) {
	fmt.Printf("%v \n", twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if v, ok := m[target-v]; ok {
			return []int{v, k}
		}
		m[v] = k
	}

	return nil
}
