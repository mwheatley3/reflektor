package reflektor

import (
	"errors"
	"testing"
)

var constructTests = []struct {
	name   string
	target interface{}
	err    error
}{
	{
		name:   "struct",
		target: struct{}{},
	},
	{
		name:   "struct pointer",
		target: &struct{}{},
	},
	{
		name: "struct pointer pointer",
		target: func() interface{} {
			v := &struct{}{}
			return &v
		}(),
	},
	{name: "int", target: 8, err: ErrInvalidTarget},
	{name: "bool", target: true, err: ErrInvalidTarget},
	{name: "func", target: func() {}, err: ErrInvalidTarget},
	{name: "chan", target: make(chan bool), err: ErrInvalidTarget},
}

func TestConstructor(t *testing.T) {
	for _, test := range constructTests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Reflekt(test.target)
			if test.err == nil {
				if err == nil {
					return
				}

				t.Fatalf("Unexpected err: %s", err)
			}

			if err == nil {
				t.Fatalf("Unexpected nil, but expected err: %s", test.err)
			}
			if errors.Is(err, test.err) {
				return
			}

			t.Fatalf("Expected %s, but found %s", test.err, err)
		})
	}
}
