package reflektor

import (
	"encoding/json"
	"fmt"
	"os"
)

type Printer interface {
	Print(r *Result) error
}

type JSONPrinter struct {
	Prefix string
	Indent string
}

func (j JSONPrinter) Print(r *Result) error {
	for _, val := range r.value {
		bytes, err := json.MarshalIndent(val.Interface(), j.Prefix, j.Indent)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, string(bytes))
	}
	return nil
}