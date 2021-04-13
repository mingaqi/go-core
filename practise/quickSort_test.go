package practise

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22, 18}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		i, j := low, high
		key := arr[(i+j)/2] // 取中间数作为基准
		fmt.Println("基准:", key)
		for i < j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if low < j {
			quickSort(arr, low, j)
		}
		if high > i {
			quickSort(arr, i, high)
		}
	}
}
