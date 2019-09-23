package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper sleeps
type Sleeper interface {
	Sleep()
}

// DefaultSleeper sleeps defaultly
type DefaultSleeper struct{}

// Sleep sleeps
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown counts down and prints Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
