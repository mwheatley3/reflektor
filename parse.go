package reflektor

import (
	"errors"
	"reflect"
)

var (
	ErrUnrecognizedParam = errors.New("Unrecognized param found")
)

func ParseIn(arg string, in reflect.Type) (reflect.Value, error) {
	switch in.Kind() {
	case reflect.String:
		return reflect.ValueOf(arg), nil
	default:
		return reflect.Value{}, ErrUnrecognizedParam
	}
}
