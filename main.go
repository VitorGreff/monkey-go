package main

import (
	"fmt"
	"os"
	"os/user"

	"interpreter/repl"
)

func main() {
	// cache the current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to monkey repl!\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
