package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns hello to the name
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
