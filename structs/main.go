package main

import "fmt"

type Point struct {
	x uint8
	y uint8
}

type Person struct {
	firstName, lastName string
	age                 uint8
}

func main() {
	var point Point = Point{10, 255}
	person := Person{
		firstName: "Paulo",
		lastName:  "Gabriel",
		age:       22,
	}
	point.x = 42
	v := &point
	v.y = 69
	fmt.Println(point)
	fmt.Println(person)
}
