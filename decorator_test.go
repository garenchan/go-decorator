package decorator

import (
	"reflect"
	"testing"
)

// div integer division
func div(a, b int) int {
	return a / b
}

// safeDiv returns 0 when the divisor is 0
func safeDiv(targetFunc reflect.Value, in []reflect.Value) (out []reflect.Value) {
	if len(in) >= 2 {
		divisor := in[1].Interface().(int)
		if divisor == 0 {
			return []reflect.Value{reflect.ValueOf(0)}
		}
	}

	return targetFunc.Call(in)
}

func TestDecorate(t *testing.T) {
	myDiv := div
	Decorate(&myDiv, div, safeDiv)

	var result int

	result = myDiv(1, 1)
	if result != 1 {
		t.Error("Expect 1, but got:", result)
	}

	result = myDiv(1, 0)
	if result != 0 {
		t.Error("Expect 0, but got:", result)
	}
}
