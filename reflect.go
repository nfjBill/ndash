package ndash

import (
	"reflect"
	"strings"
)

func ReflectIsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func ReflectElemClone(inter interface{}) interface{} {
	nInter := reflect.New(reflect.TypeOf(inter).Elem())

	val := reflect.ValueOf(inter).Elem()
	nVal := nInter.Elem()
	for i := 0; i < val.NumField(); i++ {
		nvField := nVal.Field(i)
		nvField.Set(val.Field(i))
	}

	return nInter.Interface()
}

func ReflectElemOmit(inter interface{}, inc []string) interface{} {
	v := reflect.ValueOf(inter).Elem()

	for i := 0; i < v.Type().NumField(); i++ {
		n := v.Type().Field(i).Name
		var ins bool

		for _, str := range inc {
			if strings.ToLower(n) == strings.ToLower(str) {
				ins = true
			}
		}

		if !ins {
			v.Field(i).Set(reflect.Zero(v.Field(i).Type()))
		}
	}

	return inter
}
