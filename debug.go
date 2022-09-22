package ndash

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ErrLog(err interface{}) {
	if gin.Mode() == gin.DebugMode && err != nil {
		panic(err)
		//log.Fatal(err)
	}
}

func PrintfLog(err interface{}, msg string) {
	if err != nil {
		log.Printf("[ERROR] System error: %v", msg)
	}
}

func FatalfLog(err interface{}, msg string) {
	if err != nil {
		log.Fatalf("[FATAL] System startup error: %v", msg)
	}
}

//ndash.Try(func() {
//	go RunApiV1GetCache(ctxCopy)
//}).Catch(1, func(e interface{}) {
//	log.Println("int", e)
//}).Catch("", func(e interface{}) {
//	log.Println("string", e)
//}).Finally(func() {
//	log.Println("finally")
//})

type ExceptionHandler func(e interface{})

type tryStruct struct {
	catch ExceptionHandler
	hold  func()
}

func Try(f func()) *tryStruct {
	return &tryStruct{
		catch: func(err interface{}) {},
		hold:  f,
	}
}

func (t *tryStruct) Catch(f ExceptionHandler) *tryStruct {
	t.catch = f
	return t
}

func (t *tryStruct) Finally() {
	defer func() {
		if e := recover(); nil != e {
			t.catch(e)
		}
	}()

	t.hold()
}

//// Try catches exception from f
//func Try(f func()) *tryStruct {
//	return &tryStruct{
//		catches: make(map[reflect.Type]ExeceptionHandler),
//		hold:    f,
//	}
//}
//
//// ExeceptionHandler handle exception
//type ExeceptionHandler func(interface{})
//
//type tryStruct struct {
//	catches map[reflect.Type]ExeceptionHandler
//	hold    func()
//}
//
//func (t *tryStruct) Catch(e interface{}, f ExeceptionHandler) *tryStruct {
//	t.catches[reflect.TypeOf(e)] = f
//	return t
//}
//
//func (t *tryStruct) Finally(f func()) {
//	defer func() {
//		if e := recover(); nil != e {
//			if h, ok := t.catches[reflect.TypeOf(e)]; ok {
//				h(e)
//			} else {
//				f()
//			}
//		}
//	}()
//
//	t.hold()
//}
