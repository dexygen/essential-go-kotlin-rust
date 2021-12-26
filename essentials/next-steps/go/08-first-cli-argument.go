package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	programPath := args[0]
	greetingReceiver := args[1]
	fmt.Println("Program path:" + programPath)
	fmt.Println("Hello,", greetingReceiver)
}
