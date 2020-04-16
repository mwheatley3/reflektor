package structs

import (
	"fmt"
	"reflect"
)

func PrintFields(s interface{}) {
	t := reflect.TypeOf(s)


	var fields []string
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}

	fmt.Println(fields)
}