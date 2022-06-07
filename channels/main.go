package main

func sumSlice(values []int, c chan int) {

	sum := 0

	for _, val := range values {
		sum += val
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func channelWithRange() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for n := range c {
		print(n, " ")
	}
}

func main() {
	values := []int{1, 4, 5, 2, 1, 4, 5, 10}

	c := make(chan int)
	go sumSlice(values[len(values)/2:], c)
	go sumSlice(values[:len(values)/2], c)

	x, y := <-c, <-c
	println(x, y, x+y)

	/*
	   buffered channels
	*/

	c2 := make(chan int, 3)
	c2 <- 1
	c2 <- 2
	c2 <- 3
	//c2 <- 4 //this will panic
	println(<-c2)
	println(<-c2)
	println(<-c2)

	channelWithRange()
}
