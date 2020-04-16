package structs

import (
	"fmt"
	"reflect"
)

func PrintMethods(s interface{}) {
	t := reflect.TypeOf(s)
	c := reflect.Zero(t)
	fmt.Printf("c %#+v\n\n", c)

	fmt.Printf("%#+v\n\n", t.NumMethod())

	methods := make([]string, 0)
	for i := 0; i < c.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}

	fmt.Println(methods)
}