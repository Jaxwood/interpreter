package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello %s!", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
