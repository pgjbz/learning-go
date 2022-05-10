package main

func main() {
	i, j := 42, 2701

	p := &i

	println("p is pointer to im and this address = ", p)
	*p = 21

	println("value of i changed by &p = ", i)

	println("j equals to = ", j)

	p = &j

	*p = 1999

	println("p points to j and change value to ", *p)

	/*
	   go dont have pointer aritmetic
	*/
}
