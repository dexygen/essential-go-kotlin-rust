package main

import "fmt"

func main() {
	var greeting string
	fmt.Printf("'greeting' variable, default value: \"%s\"\n", greeting)
	greeting = "Hello, World!"
	fmt.Println(greeting)
}
