package main

import (
	"fmt"
	"reflect"
)

func walk(x any, fn func(input string)) {

	numberOfValues := 0

	var getField func(int) reflect.Value

	switch val := getValue(x); val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func main() {
	fmt.Println("Hello World")
}
