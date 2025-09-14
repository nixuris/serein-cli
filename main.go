package main

import (
	"fmt"
	"os"
	"serein/cmd"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("serein version: %s\n", version)
		return
	}
	cmd.Execute()
}
