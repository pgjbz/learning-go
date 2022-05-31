package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		println("is int double is %v", v*2)
	case string:
		println("is string and vaue is", v)
	default:
		fmt.Printf("i don't know wtf is v %T\n", v)
	}
}

func main() {
	var i interface{} = "Hello"

	s := i.(string)
	println(s)

	s, ok := i.(string)
	println(s, ok)

	f, ok := i.(float64)

	println(f, ok)

	//f = i.(float64) //panic because i is a string not float64
	//println(f)
	do("Banana")
	do(42)
	do(25.0)
}
