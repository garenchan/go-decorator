# go-decorator

go-decorator is a library that facilitates the implementation of decorator pattern.

## Installation

To install go-decorator, use go get:

```shell
go get -u github.com/garenchan/go-decorator
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/garenchan/go-decorator"
)

func sum(a, b int) int {
	fmt.Printf("%d + %d\n", a, b)
	return a + b
}

func main() {
	mySum := sum

	// Use TimeIT to decorate sum function, and get a new function mySum
	decorator.Decorate(&mySum, sum, decorator.TimeIt)

	// mySum will print the cost time.
	ret := mySum(1, 2)

	// Here will print 3.
	fmt.Println(ret)
}
```

## Decorator

You can implement your own decorators for reuse,  for example:

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/garenchan/go-decorator"
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

func main() {
	myDiv := div
	decorator.Decorate(&myDiv, div, safeDiv)

	// Here will print 0.
	fmt.Println(myDiv(1, 0))

	// Here will print 1.
	fmt.Println(myDiv(1, 1))
}
```
