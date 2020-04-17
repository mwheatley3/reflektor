package reflektor

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidMethod = errors.New("Method not found")
)

type Reflekted struct {
	target reflect.Value
}

func (r *Reflekted) Method(method string) (*Func, error) {
	m := r.target.MethodByName(method)
	if (m == reflect.Value{}) {
		return nil, ErrInvalidMethod
	}

	return &Func{m, method}, nil
}

func (r *Reflekted) Methods() []*Func {
	var (
		t       = r.target.Type()
		l       = t.NumMethod()
		methods = []*Func{}
	)

	for i := 0; i < l; i++ {
		m := t.Method(i)

		// unexported
		if m.PkgPath != "" {
			continue
		}

		methods = append(methods, &Func{r.target.Method(i), m.Name})
	}

	return methods
}

type Func struct {
	fn   reflect.Value
	Name string
}

func (fn *Func) Call(args ...string) *Result {
	var (
		in       = make([]reflect.Value, len(args))
		out      []reflect.Value
		mt       = fn.fn.Type()
		variadic = mt.IsVariadic()
	)

	for i := range args {
		var (
			err error
			typ reflect.Type
		)

		if variadic && i >= mt.NumIn() {
			typ = mt.In(mt.NumIn() - 1)
		} else {
			typ = mt.In(i)
		}

		in[i], err = ParseIn(args[i], typ)
		if err != nil {
			return nil
		}
	}

	if variadic {
		out = fn.fn.CallSlice(in)
	} else {
		out = fn.fn.Call(in)
	}

	return &Result{out}
}

type Result struct {
	value []reflect.Value
}

func (r *Result) Value(i int) interface{} {
	return nil
}

func (r *Result) HasError() (error, bool) {
	return nil, false
}
