package ndash

import (
	"log"
	"reflect"
	"strings"
)

func StructForEach(s interface{}, callback func(v reflect.Value, key reflect.StructField)) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
	}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		vf := v.Field(i)
		tf := t.Field(i)
		callback(vf, tf)
	}
}

func StructForEachSafe(s interface{}, callback func(v reflect.Value, key reflect.StructField)) {
	StructForEach(s, func(iv reflect.Value, iKey reflect.StructField) {
		if !ReflectIsBlank(iv) {
			callback(iv, iKey)
		}
	})
}

func StructKeys(s interface{}) []string {
	var tmp []string
	StructForEach(s, func(v reflect.Value, key reflect.StructField) {
		tmp = append(tmp, key.Name)
	})
	return tmp
}

func StructValues(s interface{}) []interface{} {
	var tmp []interface{}
	StructForEachSafe(s, func(v reflect.Value, key reflect.StructField) {
		tmp = append(tmp, v.Interface())
	})
	return tmp
}

//获取结构体中字段的名称
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

//获取结构体中Tag的值，如果没有tag则返回字段值
func GetTagName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}