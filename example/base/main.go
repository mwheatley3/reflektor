package main

import (
	"flag"
	"fmt"
	"github.com/jonasi/reflektor"
	"log"
	"os"
)

type Example struct {}

func (e Example) Method1(echo string) string {
	fmt.Printf("Calling method 1... echo %s", echo)
	return echo
}

func (e Example) Method2(echo string) string {
	fmt.Printf("Calling method 2... echo %s", echo)
	return echo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("name of sruct is required")
		fmt.Println("list or call subcommand is required")
		os.Exit(1)
	}

	c := ReflektedCmd{
		StructToReflekt: Example{},
		Name: "example",
	}

	if os.Args[1] != c.Name {
		fmt.Printf("Struct is not registered: %s", os.Args[1])
	}

	switch os.Args[2] {
	case "list":
		listCmd(c)
	case "call":
		callCmd(c, os.Args[3:]...)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func listCmd(st ReflektedCmd) {
	r, err := reflektor.Reflekt(st.StructToReflekt)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	funcs := r.Methods()

	for _, m := range funcs {
		fmt.Println(m.Name)
	}
}

func callCmd(st ReflektedCmd, args... string) {
	var m string
	cmd := flag.NewFlagSet("call", flag.ExitOnError)
	cmd.StringVar(&m, "method", "", "Name of Method to Call. (Required)")
	err := cmd.Parse(args)
	if err != nil {
		printfUsageExit(cmd, "error parsing command: %s\n", err)
	}

	if m == "" {
		cmd.PrintDefaults()
		os.Exit(1)
	}

	r, err := reflektor.Reflekt(st.StructToReflekt)
	if err != nil {
		log.Fatalf("err %v", err)
	}
	f, err := r.Method(m)
	if err != nil {
		log.Fatalf("err %v", err)
	}

	result := f.Call(args[2:]...)
	fmt.Printf("result%#+v\n\n", result)
	// TODO print results
}

func printfUsageExit(cmd *flag.FlagSet, fmtMsg string, argv ...interface{}) {
	fmt.Fprintf(os.Stderr, fmtMsg, argv...)
	fmt.Printf("usage: relektor %s\n", cmd.Name())
	cmd.PrintDefaults()
	os.Exit(1)
}

type ReflektedCmd struct {
	StructToReflekt interface{}
	Name string
}