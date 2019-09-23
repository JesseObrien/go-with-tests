package main 

import (
	"fmt"
	"io"
)

// Greet allows us to greet a user
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
