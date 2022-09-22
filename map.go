package ndash

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type IGetMapValue interface {
	//LastError() error
	//MustBeValid() IAny
	//ToBool() bool
	//ToInt() int
	//ToInt32() int32
	//ToInt64() int64
	//ToUint() uint
	//ToUint32() uint32
	//ToUint64() uint64
	//ToFloat32() float32
	//ToFloat64() float64
	ToString() string
	//ToVal(val interface{})
	//Get(path ...interface{}) IAny
	//Size() int
	//Keys() []string
	GetInterface() interface{}
}

type SGetMapValue struct {
	val interface{}
	key string
}

func GetMapValue(val interface{}, key string) IGetMapValue {
	return &SGetMapValue{val, key}
}

func (any *SGetMapValue) GetInterface() interface{} {
	return any.val.(map[string]interface{})[any.key]
}

func (any *SGetMapValue) ToString() string {
	src := any.val.(map[string]interface{})[any.key]
	if src == nil {
		return ""
	}
	return src.(string)
}

const (
	OverwriteExistingDstKey = 1 << iota
	ErrorOnExistingDstKey
	ErrorOnDifferentDstKeyValue
)

func MergeInto(dst, src interface{}, flags int) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	if dstVal.Kind() != reflect.Map {
		return fmt.Errorf("dst is not a valid map: %v", dstVal.Kind())
	}
	if srcVal.Kind() != reflect.Map {
		return fmt.Errorf("src is not a valid map: %v", srcVal.Kind())
	}
	if dstTyp, srcTyp := dstVal.Type(), srcVal.Type(); !dstTyp.AssignableTo(srcTyp) {
		return fmt.Errorf("type mismatch, can't assign '%v' to '%v'", srcTyp, dstTyp)
	}

	if dstVal.IsNil() {
		return fmt.Errorf("dst value is nil")
	}
	if srcVal.IsNil() {
		// Nothing to merge
		return nil
	}

	for _, k := range srcVal.MapKeys() {
		if dstVal.MapIndex(k).IsValid() {
			if flags&ErrorOnExistingDstKey != 0 {
				return fmt.Errorf("dst key already set (ErrorOnExistingDstKey=1), '%v'='%v'", k, dstVal.MapIndex(k))
			}
			if dstVal.MapIndex(k).String() != srcVal.MapIndex(k).String() {
				if flags&ErrorOnDifferentDstKeyValue != 0 {
					return fmt.Errorf("dst key already set to a different value (ErrorOnDifferentDstKeyValue=1), '%v'='%v'", k, dstVal.MapIndex(k))
				}
				if flags&OverwriteExistingDstKey != 0 {
					dstVal.SetMapIndex(k, srcVal.MapIndex(k))
				}
			}
		} else {
			dstVal.SetMapIndex(k, srcVal.MapIndex(k))
		}
	}

	return nil
}

func MapEachSort(eachMap interface{}, eachFunc interface{}) {
	eachMapValue := reflect.ValueOf(eachMap)
	eachFuncValue := reflect.ValueOf(eachFunc)
	eachMapType := eachMapValue.Type()
	eachFuncType := eachFuncValue.Type()
	if eachMapValue.Kind() != reflect.Map {
		panic(errors.New("ksort.EachMap failed. parameter \"eachMap\" type must is map[...]...{}"))
	}
	if eachFuncValue.Kind() != reflect.Func {
		panic(errors.New("ksort.EachMap failed. parameter \"eachFunc\" type must is func(key ..., value ...)"))
	}
	if eachFuncType.NumIn() != 2 {
		panic(errors.New("ksort.EachMap failed. \"eachFunc\" input parameter count must is 2"))
	}
	if eachFuncType.In(0).Kind() != eachMapType.Key().Kind() {
		panic(errors.New("ksort.EachMap failed. \"eachFunc\" input parameter 1 type not equal of \"eachMap\" key"))
	}
	if eachFuncType.In(1).Kind() != eachMapType.Elem().Kind() {
		panic(errors.New("ksort.EachMap failed. \"eachFunc\" input parameter 2 type not equal of \"eachMap\" value"))
	}

	// 对key进行排序
	// 获取排序后map的key和value，作为参数调用eachFunc即可
	switch eachMapType.Key().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		keys := make([]int, 0)
		keysMap := map[int]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, int(value.Int()))
			keysMap[int(value.Int())] = value
		}
		sort.Ints(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.Float64, reflect.Float32:
		keys := make([]float64, 0)
		keysMap := map[float64]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, float64(value.Float()))
			keysMap[float64(value.Float())] = value
		}
		sort.Float64s(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.String:
		keys := make([]string, 0)
		keysMap := map[string]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, value.String())
			keysMap[value.String()] = value
		}
		sort.Strings(keys)
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	default:
		panic(errors.New("\"eachMap\" key type must is int or float or string"))
	}
}

//func MapMerge(src interface{}, dst ...interface{}) interface{} {
//	var m = make(map[string]interface{})
//	s := src.(map[string]interface{})
//	for k, v := range s {
//		m[k] = v
//	}
//
//	for _, ds := range dst {
//		d := ds.(map[string]interface{})
//		for k, v := range d {
//			m[k] = v
//		}
//	}
//
//	return m
//}
