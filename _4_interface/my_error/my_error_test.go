package my_error

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var e interface{}
	e = New(401, "an error")
	switch e.(type) {
	case error:
		fmt.Println("this is an error type")
	}

	fmt.Println(New(401, "an error"))
}
