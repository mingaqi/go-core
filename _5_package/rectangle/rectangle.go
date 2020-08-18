package rectangle

import (
	"fmt"
	"math"
)

func Area(l, w float64) float64 {
	area := l * w
	return area
}

func Diagonal(l, w float64) float64 {
	res := math.Sqrt((l * l) + (w * w))
	return res
}

func init() {
	fmt.Println("rectangle package initialized")
}

func init() {
	fmt.Println("aaaa")
}
