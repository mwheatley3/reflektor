package reflektor

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

var parseTests = []struct {
	name     string
	str      []string
	expected interface{}
	err      error
}{
	{name: "string", str: []string{"string"}, expected: "string"},
	{name: "empty", str: []string{""}, expected: ""},
	{name: "*string", str: []string{"string"}, expected: func() *string { x := "string"; return &x }()},
	{name: "bool - true", str: []string{"true", "1"}, expected: true},
	{name: "bool - false", str: []string{"false", "0"}, expected: false},
	{name: "bool - bad syntax", str: []string{"", "no", "Y", "SDFJSLDF JSLKDFJ SDKFJ SLDKFJLKDFJSD"}, expected: false, err: strconv.ErrSyntax},
	{name: "int - zero", str: []string{"0"}, expected: 0},
	{name: "int - positive", str: []string{"100"}, expected: 100},
	{name: "int - negative", str: []string{"-100"}, expected: -100},
	{name: "int8 ", str: []string{"100"}, expected: int8(100)},
	{name: "int16", str: []string{"100"}, expected: int16(100)},
	{name: "int32", str: []string{"100"}, expected: int32(100)},
	{name: "int64", str: []string{"100"}, expected: int64(100)},
	{name: "int8 - overflow", str: []string{"512"}, err: strconv.ErrRange, expected: int8(0)},
	{name: "uint", str: []string{"100"}, expected: uint(100)},
	{name: "uint8", str: []string{"100"}, expected: uint8(100)},
	{name: "uint16", str: []string{"100"}, expected: uint16(100)},
	{name: "uint32", str: []string{"100"}, expected: uint32(100)},
	{name: "uint64", str: []string{"100"}, expected: uint64(100)},
	{name: "float32", str: []string{"100.01"}, expected: float32(100.01)},
	{name: "float64", str: []string{"100.01"}, expected: float64(100.01)},
	{name: "channel", str: []string{"100.01"}, err: ErrUnrecognizedParam, expected: make(chan bool)},
	{name: "func", str: []string{"100.01"}, err: ErrUnrecognizedParam, expected: func() {}},
	{name: "slice - int", str: []string{"[1,2]"},  expected: []int{1,2}},
	{name: "slice - string", str: []string{"[abc,hello]"},  expected: []string{"abc","hello"}},
	{name: "slice - float", str: []string{"[1.1,3.4]"},  expected: []float64{1.1,3.4}},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		suffix := len(test.str) > 1
		for _, str := range test.str {
			name := test.name
			if suffix {
				name += " " + str
			}

			t.Run(name, func(t *testing.T) {
				dest := reflect.New(reflect.TypeOf(test.expected)).Elem()
				err := ParseIn(str, dest)
				if err != nil {
					if errors.Is(err, test.err) {
						return
					}

					if test.err == nil {
						t.Fatalf("Unexpected error: %s", err)
					}

					t.Fatalf("Expected error %s but found %s", test.err, err)
				}

				if reflect.DeepEqual(dest.Interface(), test.expected) {
					return
				}

				t.Fatalf("Expected value %#v, but found %#v", test.expected, dest.Interface())
			})
		}
	}
}
