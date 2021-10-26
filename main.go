package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s - Asimov 0.1 command prompt\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
