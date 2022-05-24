package main

import "math"

type MyFloat float64

/*
  cannot declare methods in types declared in others packages

*/

func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

func main() {
	v := MyFloat(-math.Sqrt2)
	println(v.Abs())
}
