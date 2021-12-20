package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var greeting = "Greetings"
	var greetingReceiver string
	var scanner = bufio.NewScanner(os.Stdin)

	for greetingReceiver == "" {
		fmt.Print("Who are you greeting? ")
		scanner.Scan()
		greetingReceiver = strings.TrimSpace(scanner.Text())
	}

	fmt.Println(greeting, greetingReceiver)
}
