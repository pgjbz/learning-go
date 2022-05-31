package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int8
}

func (p Person) String() string {
	return fmt.Sprintf("name := %v, age := %v", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (m MyError) Error() string {
	return fmt.Sprintf("at %v, %s", m.When, m.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"dont work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	p := Person{Name: "Paulo", Age: 22}
	fmt.Println(p)
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
