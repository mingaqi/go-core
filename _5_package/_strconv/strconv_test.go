package _strconv

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {
	str := "100000"
	//i, _ := strconv.ParseInt(str, 10, 32)
	i, _ := strconv.Atoi(str)
	str = strconv.Itoa(i)
	fmt.Println(i)

	b := "false"
	parseBool, _ := strconv.ParseBool(b)
	fmt.Println(parseBool)

	f := "3.323"
	float, _ := strconv.ParseFloat(f, 64)
	fmt.Println(float)
}
