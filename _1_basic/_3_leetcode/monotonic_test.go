package _3_leetcode

import "testing"

// 单调数组
// 可优化为一次循环
func TestName(t *testing.T) {
	i := &[]int{5, 3, 2, 4, 1}
	println(isMonotonic(*i))
}

func isMonotonic(A []int) bool {

	size := len(A)
	if size <= 1 {
		return true
	}

	flag := false
	flag2 := false
	for k, v := range A {
		if k+1 < size && v <= A[k+1] {
			if k+2 == size {
				flag = true
				break
			}
			flag = true
			continue
		} else {
			flag = false
			break
		}
	}

	for k, v := range A {
		if k+1 < size && v >= A[k+1] {
			if k+2 == size {
				flag2 = true
				break
			}
			flag2 = true
			continue
		} else {
			flag2 = false
			break
		}
	}
	if flag || flag2 {
		return true
	} else {
		return false
	}

}
