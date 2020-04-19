package reflektor

import (
	"log"
	"testing"
)

type ExamplePtrReceiver struct {
	Field1 string
	Field2 string
}

func (e *ExamplePtrReceiver) Method1() {
	println("method1 called")
}

func (e *ExamplePtrReceiver) Method2() {
	println("method2 called")
}

func (e *ExamplePtrReceiver) Method3() {
	println("method3 called")
}

type Example struct {
	Field1 string
	Field2 string
}

func (e Example) Method1() {
	println("method1 called")
}

func (e Example) Method2() {
	println("method2 called")
}

func (e Example) Method3() {
	println("method3 called")
}

func TestMethods(t *testing.T) {
	want := map[string]bool{
		"Method1":true,
		"Method2":true,
		"Method3":true,
	}

	sPtr := ExamplePtrReceiver{}
	rPtr, err := Reflekt(&sPtr)
	if err != nil {
		log.Fatal(err)
	}

	ptrFuncs := rPtr.Methods()
	if len(ptrFuncs) != len(want) {
		log.Fatalf("Wrong number of ptrFuncs, want: %d, got: %d", len(want), len(ptrFuncs))
	}
	for _, m := range ptrFuncs {
		if _, ok := want[m.Name]; !ok {
			log.Fatalf("Function %s not found in want map", m.Name)
		}
	}

	s := Example{}
	r, err := Reflekt(s)
	if err != nil {
		log.Fatal(err)
	}

	funcs := r.Methods()
	if len(funcs) != len(want) {
		log.Fatalf("Wrong number of funcs, want: %d, got: %d", len(want), len(funcs))
	}
	for _, m := range funcs {
		if _, ok := want[m.Name]; !ok {
			log.Fatalf("Function %s not found in want map", m.Name)
		}
	}
}


