package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}

func Perimeter(retangle Rectangle) float64 {
	return (retangle.Height + retangle.Width) * 2
}

func main() {
	fmt.Println("Hello World")
}
