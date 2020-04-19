package reflektor

import (
	"reflect"
	"testing"
)

type empty struct{}

type example struct {
	Field1 string
	Field2 string
}

func (e *example) Method1()     {}
func (e *example) Method2()     {}
func (e *example) Method3()     {}
func (e example) ValueMethod1() {}
func (e *example) unexported()  {}

var listTests = []struct {
	name    string
	target  interface{}
	methods []string
}{
	{
		name:    "empty struct",
		target:  struct{}{},
		methods: []string{},
	},
	{
		name:    "ptr to struct",
		target:  &example{},
		methods: []string{"Method1", "Method2", "Method3", "ValueMethod1"},
	},
	{
		name:    "struct",
		target:  example{},
		methods: []string{"ValueMethod1"},
	},
}

func TestListMethods(t *testing.T) {
	for _, test := range listTests {
		t.Run(test.name, func(t *testing.T) {
			r, err := Reflekt(test.target)
			if err != nil {
				t.Fatalf("error calling Reflekt: %s", err)
			}

			ms := []string{}
			for _, m := range r.Methods() {
				ms = append(ms, m.Name)
			}

			if !reflect.DeepEqual(test.methods, ms) {
				t.Fatalf("expected methods %s but found %s", test.methods, ms)
			}
		})
	}
}
