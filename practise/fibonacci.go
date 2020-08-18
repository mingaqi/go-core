package practise

const loops int = 51

var fibos   [loops]uint64

/*基于缓存的斐波那契数列*/
func Fibonacci(i int) (r uint64) {
	if fibos[i] != 0 {
		r = fibos[i]
		return
	}

	if i <= 1 {
		r = 1
	} else {
		r = Fibonacci(i-1) + Fibonacci(i-2)
	}
	fibos[i] = r
	return
}
