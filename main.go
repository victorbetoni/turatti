package main

import (
	"fmt"
	"os"
	"turatti/repl"
)

func main() {
	fmt.Printf("Running Turatti lang REPL...\n")
	repl.Start(os.Stdin, os.Stdout)
}
