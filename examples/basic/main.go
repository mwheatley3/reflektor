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

func (e *example) Struct() Stuff {
	return Stuff{
		Strings{"a", "b", "c"},
		Numbers{1,2,3},
	}
}

type Stuff struct {
	Strings Strings
	Numbers Numbers
}

type Strings struct {
	A string
	B string
	C string
}

type Numbers struct {
	One int
	Two int
	Three int
}

func main() {
	reflektor.Command(&example{})
}
