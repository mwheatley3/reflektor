package reflektor

import (
	"errors"
	"math/bits"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrUnrecognizedParam = errors.New("Unrecognized param found")
)

func ParseIn(arg string, dest reflect.Value) error {
	switch dest.Kind() {
	case reflect.String:
		dest.SetString(arg)
	case reflect.Bool:
		v, err := strconv.ParseBool(arg)
		if err != nil {
			return err
		}

		dest.SetBool(v)
	case reflect.Ptr:
		if dest.IsNil() {
			dest.Set(reflect.New(dest.Type().Elem()))
		}
		return ParseIn(arg, dest.Elem())
	case reflect.Int:
		if err := parseInt(arg, bits.UintSize, dest); err != nil {
			return err
		}
	case reflect.Int8:
		if err := parseInt(arg, 8, dest); err != nil {
			return err
		}
	case reflect.Int16:
		if err := parseInt(arg, 16, dest); err != nil {
			return err
		}
	case reflect.Int32:
		if err := parseInt(arg, 32, dest); err != nil {
			return err
		}
	case reflect.Int64:
		if err := parseInt(arg, 64, dest); err != nil {
			return err
		}
	case reflect.Uint:
		if err := parseUint(arg, bits.UintSize, dest); err != nil {
			return err
		}
	case reflect.Uint8:
		if err := parseUint(arg, 8, dest); err != nil {
			return err
		}
	case reflect.Uint16:
		if err := parseUint(arg, 16, dest); err != nil {
			return err
		}
	case reflect.Uint32:
		if err := parseUint(arg, 32, dest); err != nil {
			return err
		}
	case reflect.Uint64:
		if err := parseUint(arg, 64, dest); err != nil {
			return err
		}
	case reflect.Float32:
		if err := parseFloat(arg, 32, dest); err != nil {
			return err
		}
	case reflect.Float64:
		if err := parseFloat(arg, 64, dest); err != nil {
			return err
		}
	case reflect.Slice:
		if err := parseSlice(arg, dest); err != nil {
			return err
		}
	default:
		return ErrUnrecognizedParam
	}

	return nil
}

func parseInt(arg string, size int, dest reflect.Value) error {
	v, err := strconv.ParseInt(arg, 10, size)
	if err != nil {
		return err
	}

	dest.SetInt(v)
	return nil
}

func parseUint(arg string, size int, dest reflect.Value) error {
	v, err := strconv.ParseUint(arg, 10, size)
	if err != nil {
		return err
	}

	dest.SetUint(v)
	return nil
}

func parseFloat(arg string, size int, dest reflect.Value) error {
	v, err := strconv.ParseFloat(arg, size)
	if err != nil {
		return err
	}

	dest.SetFloat(v)
	return nil
}

func parseSlice(arg string, dest reflect.Value) error {
	newSlice := strings.Split(arg[1:len(arg)-1], ",")
	slice := reflect.MakeSlice(dest.Type(), len(newSlice), len(newSlice))
	for i := 0; i < len(newSlice); i++ {
		v := slice.Index(i)
		err := ParseIn(newSlice[i], v)
		if err != nil {
			return err
		}
	}

	dest.Set(slice)
	return nil
}
