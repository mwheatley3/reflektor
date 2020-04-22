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

	return newFunc(m, method), nil
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

		f := newFunc(r.target.Method(i), m.Name)
		methods = append(methods, f)
	}

	return methods
}

func newFunc(fn reflect.Value, name string) *Func {
	f := &Func{
		fn:   fn,
		Name: name,
	}
	f.args()
	return f
}

type Func struct {
	fn       reflect.Value
	Name     string
	InTypes  []string
	OutTypes []string
}

func (fn *Func) Call(args ...string) *Result {
	var (
		in       = make([]reflect.Value, len(args))
		out      []reflect.Value
		mt       = fn.fn.Type()
		variadic = mt.IsVariadic()
	)

	for i := range args {
		var typ reflect.Type

		if variadic && i >= mt.NumIn() {
			typ = mt.In(mt.NumIn() - 1)
		} else {
			typ = mt.In(i)
		}

		in[i] = reflect.New(typ).Elem()
		if err := ParseIn(args[i], in[i]); err != nil {
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

func (fn *Func) args() {
	var (
		mt     = fn.fn.Type()
		numIn  = mt.NumIn()
		numOut = mt.NumOut()
	)

	for i := 0; i < numIn; i++ {
		inValue := mt.In(i)
		k := inValue.Kind()
		fn.InTypes = append(fn.InTypes, k.String())
	}

	for i := 0; i < numOut; i++ {
		outValue := mt.Out(i)
		k := outValue.Kind()
		fn.OutTypes = append(fn.OutTypes, k.String())
	}
}

type Result struct {
	value []reflect.Value
}

func (r *Result) Len() int {
	return len(r.value)
}

func (r *Result) Value(i int) interface{} {
	return r.value[i].Interface()
}

func (r *Result) HasError() (error, bool) {
	return nil, false
}
