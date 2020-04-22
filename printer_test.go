package reflektor

import (
	"fmt"
	"testing"
)

type rv struct {
	Field1 string
	Field2 string
}

type Stuff struct {
	A string
	B string
	C string
}

func (e *rv) Method1() Stuff { return Stuff{"a", "b", "c"}}
//func (e *rv) Method2() stuff {}
//func (e *rv) Method3() stuff {}

var printTests = []struct {
	name    string
	target  	interface{}
	expected string
}{
	{
		name:    "empty struct",
		target: &rv{},
		expected: ``,
	},
	//{
	//	name:    "ptr to struct",
	//	target:  &example{},
	//	expected: []string{"Method1", "Method2", "Method3", "ValueMethod1"},
	//},
	//{
	//	name:    "struct",
	//	target:  example{},
	//	expected: []string{"ValueMethod1"},
	//},
}

func TestPrinter(t *testing.T) {
	for _, test := range printTests {
		t.Run(test.name, func(t *testing.T) {
			r, err := Reflekt(test.target)
			if err != nil {
				t.Fatalf("error calling Reflekt: %s", err)
			}

			m, err := r.Method("Method1")
			if err != nil {
				t.Fatalf("error getting method: %s", err)
			}

			result := m.Call()
			p := JSONPrinter{}
			fmt.Printf("\n\nrv%#+v\n\n", result.value[0])
			p.Print(result)

			//if !reflect.DeepEqual(test.methods, ms) {
			//	t.Fatalf("expected methods %s but found %s", test.methods, ms)
			//}
		})
	}
}