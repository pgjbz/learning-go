package countdown

import (
	"fmt"
	"io"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func NewConfigurableSleeper(duration time.Duration,
	sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
