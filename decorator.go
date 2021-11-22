package decorator

import (
	"fmt"
	"reflect"
	"time"
)

// Decorator is a decorator function.
type Decorator func(targetFunc reflect.Value, in []reflect.Value) (out []reflect.Value)

// Decorate decorate target with multiple decorators, and return with decoratedPtr.
func Decorate(decoratedPtr, target interface{}, decors ...Decorator) (err error) {
	decoratedFunc := reflect.ValueOf(decoratedPtr).Elem()
	targetFunc := reflect.ValueOf(target)
	currentFunc := targetFunc

	for _, decor := range decors {
		currentFunc = makeFunc(currentFunc, decor)
	}

	decoratedFunc.Set(currentFunc)
	return
}

// makeFunc decorate targetFunc with decor, and return decorated func.
func makeFunc(targetFunc reflect.Value, decor Decorator) reflect.Value {
	val := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			return decor(targetFunc, in)
		})

	return val
}

// TimeIt measure the execution time of the function and print.
func TimeIt(targetFunc reflect.Value, in []reflect.Value) (out []reflect.Value) {
	defer func(t time.Time) {
		fmt.Printf("Cost %v\n", time.Since(t))
	}(time.Now())

	out = targetFunc.Call(in)
	return
}
