package main

import (
	"fmt"
	"time"
)

func say(times int) {
	for i := 0; i < times; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	go say(10) //execute in another thread WOW
	say(10)
}
