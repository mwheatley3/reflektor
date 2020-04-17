package reflektor

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidTarget = errors.New("Invalid reflekt target. Must be a struct, interface or pointer to one.")
)

func Reflekt(obj interface{}) (*Reflekted, error) {
	rv := unwrap(reflect.ValueOf(obj))
	if !canReflekt(rv) {
		return nil, ErrInvalidTarget
	}

	return &Reflekted{target: rv}, nil
}

func canReflekt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Struct, reflect.Interface:
		return true
	default:
		return false
	}
}

func unwrap(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}
