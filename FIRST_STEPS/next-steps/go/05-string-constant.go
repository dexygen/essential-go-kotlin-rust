package main

import "fmt"

func main() {
	// const greeting string // compiler error
	const greeting string = "Hello, World!"
	// greeting = "Greetings, Earthling" // compiler error
	fmt.Println(greeting)
}
