package main

import "time"

func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		/*
		   select -> wait for channel receive a message,
		   when any channel receive a message execute the case,
		   if more than one receive a message is random selected
		*/
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}
}

func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			println(<-c) //consumes C
		}
		quit <- 0
	}()
	fibo(c, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			println("tick.")
		case <-boom:
			println("BOOM!")
			return
		default: //default always execute if no message
			println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
