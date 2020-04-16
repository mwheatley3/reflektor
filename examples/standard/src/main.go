package main

import (
	"github.com/jonasi/reflektor/examples/standard/src/structs"
	"os"
)

func main() {
	args := os.Args[1:]
	if args[0] == "list" {
		structs.PrintFields(structs.Example{})
		structs.PrintMethods(structs.Example{})
	}
}
