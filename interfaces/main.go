package main

import (
	"fmt"
	"math"
)

type Trigonometry interface {
	Angle(x, y float64) float64
}

type I interface {
	M()
}

type T struct {
	message string
}

type A struct {
	value int8
}

func (a A) M() {
	println(a.value)
}

/*
  in go don't have implements keywords
  interface are implicit implemented
*/
func (t *T) M() { //remove * in '*T' and code would not compile, because &T implements interface, not T
	if t == nil {
		return
	}
	println(t.message)
}

type Point struct {
	x, y float64
}

func (p Point) Angle(x, y float64) float64 {
	return math.Atan2(p.y-y, p.x-x) * (180 / math.Pi)
}

func main() {
	p := Point{10.0, 5.0}
	println(p.Angle(4.0, 20.2))

	var x *Point = &p           //*Point does not implement the interface Trigonometry but, Go compiler deref the value and call the function
	println(x.Angle(4.0, 20.2)) //like same (*v).Angle(4.0, 20.2)

	var i I = &T{"Hello"}
	describe(i)
	i.M()

	var t *T
	i = t
	i.M() //panic if not "workarout"

	i = A{value: 10} //go have polymorphism
	describe(i)
	i.M()

	/*
	   if call a method in nil interface will cause a runtime error
	   because has no concret type
	*/

	var b interface{} //like a Object in Java, accept any type, because any type implement zero or more methods
	b = 42
	describe(b)

	b = "Hello"
	describe(b)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
