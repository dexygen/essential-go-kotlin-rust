package main

import "fmt"

func main() {
	var greeting string = "Hello, World!"

	fmt.Printf("'greeting' variable: type: %T\n", greeting)
	// greeting = 1 // compiler error

	fmt.Println(greeting)
}
