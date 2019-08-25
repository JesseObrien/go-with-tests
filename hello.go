package main

import "fmt"

// Hello returns hello to the name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
