package main

import "fmt"

/*
  A closure function is a function that access variables out of your scope
  in this case each adder call return a function that have your wown sum value
  that access sum individualy
*/
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(i*-2),
		)
	}
}
