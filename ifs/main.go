package main

import "math"

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	i := 10
	if i == 10 {
		println("uau")
	}

	println(pow(3, 2, 10), pow(3, 3, 10))
}
