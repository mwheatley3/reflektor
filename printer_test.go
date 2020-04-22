package reflektor

import (
	"io/ioutil"
	"os"
	"testing"
)

type pExample struct {
	Field1 string
	Field2 string
}

type Stuff struct {
	A string
	B string
	C string
}

func (e *pExample) Method1() Stuff { return Stuff{"a", "b", "c"}}
func (e *pExample) Method2() string { return "HEY"}

var printTests = []struct {
	name    string
	target  	interface{}
	printer Printer
	method string
	expected string
}{
	{
		name:    "print struct",
		target: &pExample{},
		printer: JSONPrinter{
			Prefix: "",
			Indent: "  ",
		},
		method: "Method1",
		expected: `{
  "A": "a",
  "B": "b",
  "C": "c"
}`},
		{
			name:    "print string",
			target: &pExample{},
			printer: JSONPrinter{
				Prefix: "",
				Indent: "  ",
			},
			method: "Method2",
			expected: `"HEY"`,
	},
}

func TestPrinter(t *testing.T) {
	for _, test := range printTests {
		t.Run(test.name, func(t *testing.T) {
			r, err := Reflekt(test.target)
			if err != nil {
				t.Fatalf("error calling Reflekt: %s", err)
			}

			m, err := r.Method(test.method)
			if err != nil {
				t.Fatalf("error getting method: %s", err)
			}

			rescueStdout := os.Stdout
			rd, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("error creating pipe: %s", err)
			}
			os.Stdout = w

			result := m.Call()
			err = test.printer.Print(result)
			if err != nil {
				t.Fatalf("error printing: %s", err)
			}

			w.Close()
			out, err := ioutil.ReadAll(rd)
			if err != nil {
				t.Fatalf("error reading: %s", err)
			}
			os.Stdout = rescueStdout

			if string(out) != test.expected {
				t.Fatalf("Sorry!: %s", err)
			}
		})
	}
}