package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	nameArg := os.Args[1:]
	name := strings.Join(nameArg, " ")

	if len(name) < 1{
		fmt.Println("Who are you?")
		return
	}

	fmt.Printf("Hello %s!\n", name)
}
