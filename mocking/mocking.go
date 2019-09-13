package mocking

import (
	"fmt"
	"io"
	"os"
)

// Countdown counts down and prints Go!
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
