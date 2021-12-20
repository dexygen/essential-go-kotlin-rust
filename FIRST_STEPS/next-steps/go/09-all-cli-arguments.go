package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		cliArgs := os.Args[1:]
		fmt.Println("Hello,", strings.Join(cliArgs, " "))
	} else {
		fmt.Println("Please specify at least one command line argument")
	}
}
