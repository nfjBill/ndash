package ndash

import (
	"fmt"
	"reflect"
)

func InterfaceIsNil(s interface{}) bool {
	if s == nil {
		//fmt.Printf("judge1:s(%v) == nil\n", s)
		return true
	}
	if reflect.ValueOf(s).IsNil() {
		//fmt.Printf("judge2:s(%v) reflect nil\n", s)
		return true
	}
	fmt.Printf("s(%v) pass\n", s)
	return false
}
