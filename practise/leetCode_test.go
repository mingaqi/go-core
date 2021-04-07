package practise

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var number int
	var temp string = fmt.Sprintf("%d", number)
	if len(temp) < 3 {
		fmt.Println(temp)
	}
}
