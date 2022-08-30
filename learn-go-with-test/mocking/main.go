package main

import (
	"os"
	"time"

	"pgjbz.dev/mocking/countdown"
)



func main() {
	sleeper := countdown.NewConfigurableSleeper(1 * time.Second, time.Sleep)
	countdown.Countdown(os.Stdout, sleeper)
}