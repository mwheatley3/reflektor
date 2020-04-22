package main

import "github.com/jonasi/reflektor"

type example struct {
}

func (e *example) String() string {
	return "HEY"
}

func (e *example) Echo(str string) string {
	return str
}

func main() {
	reflektor.Command(&example{})
}
