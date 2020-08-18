package practise

import "testing"

func TestFibonacci(t *testing.T) {
	var i int
	for i = 0; i < loops; i++ {
		print(Fibonacci(i), ",")
	}
}
