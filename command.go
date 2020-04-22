package reflektor

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Command(obj interface{}) {
	if err := cmd(obj); err != nil {
		fmt.Fprintf(os.Stderr, "Reflektor error: %s\n", err)
		cmdUsage()
		os.Exit(1)
	}
}

func cmd(obj interface{}) error {
	args := os.Args[1:]
	if len(args) == 0 {
		return errors.New("No subcommand found")
	}

	switch strings.ToLower(args[0]) {
	case "list":
		return cmdList(obj)
	case "call":
		return cmdCall(obj, args[1:])
	default:
		return fmt.Errorf("Invalid subcommand %s", args[0])
	}
}

func cmdList(obj interface{}) error {
	r, err := Reflekt(obj)
	if err != nil {
		return err
	}

	for _, m := range r.Methods() {
		fmt.Fprintln(os.Stdout, cmdSig(m))
	}

	return nil
}

func cmdSig(f *Func) string {
	in := strings.Join(f.InTypes, ", ")
	out := strings.Join(f.OutTypes, ", ")
	if len(f.OutTypes) > 1 {
		out = "(" + out + ")"
	}
	if len(out) > 0 {
		out = " " + out
	}

	return fmt.Sprintf("%s(%s)%s", f.Name, in, out)
}

func cmdCall(obj interface{}, args []string) error {
	if len(args) == 0 {
		return errors.New("No method passed into call")
	}

	r, err := Reflekt(obj)
	if err != nil {
		return err
	}

	m, err := r.Method(args[0])
	if err != nil {
		return err
	}

	args = args[1:]
	res := m.Call(args...)

	p := JSONPrinter{
		Indent: "  ",
	}
	print(p, res)

	return nil
}

func print(p Printer, r *Result) error {
	return p.Print(r)
}

func cmdUsage() {
	name := os.Args[0]
	fmt.Fprintf(os.Stderr, `Usage:
  %s list
  %s call <method> <argument>...
`, name, name)
}
