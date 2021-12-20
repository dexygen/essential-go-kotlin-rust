package main

import "fmt"

func main() {
	var greeting = "Hello"
	greeting = greeting + ","
	greeting += " World!"
	fmt.Println(greeting)
}
