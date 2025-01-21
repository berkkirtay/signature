package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("No command is specified.")
	}
	handleCommand(os.Args[1], os.Args[2:]...)
}
