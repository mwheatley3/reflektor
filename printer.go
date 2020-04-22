package reflektor

import (
	"encoding/json"
	"fmt"
)

type Printer interface {
	Print(r *Result) error
}

type JSONPrinter struct {}

func (j JSONPrinter) Print(r *Result) error {
	for _, val := range r.value {
		println("\nhello\n")

		//for i := 0; i < val.NumField(); i++ {
		//	valueField := val.Field(i)
		//	f := valueField.Interface()
			bytes, err := json.MarshalIndent(val.Interface(), "", "")
			if err != nil {
				return err
			}
			//fmt.Fprintf(os.Stdout, string(bytes))
			fmt.Printf("bytes%#+v\n\n", string(bytes))
		//}

		println("\ngoodbye\n")
	}
	return nil
}

//func display(v reflect.Value) string {
//	switch v.Kind() {
//	case reflect.Invalid:
//		return "invalid"
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		return strconv.FormatInt(v.Int(), 10)
//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//		return strconv.FormatUint(v.Uint(), 10)
//	case reflect.Float32:
//		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
//	case reflect.Float64:
//		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
//	case reflect.Bool:
//		if v.Bool() {
//			return "true"
//		}
//		return "false"
//	case reflect.String:
//		return strconv.Quote(v.String())
//	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
//		return v.Type().String() + " 0x" +
//			strconv.FormatUint(uint64(v.Pointer()), 16)
//	default: // reflect.Array, reflect.Struct, reflect.Interface
//		return v.Type().String() + " value"
//	}
//}