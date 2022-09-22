package ndash

import (
	"bytes"
	"encoding/gob"
	"github.com/axgle/mahonia"
	"github.com/ginkgoch/godash"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

type SMsqStructToMapOptions struct {
	StrAllowEmpty bool
	StrAllowEmptyKey godash.DashSlice
}

func StrToInt(str string) int {
	re, err := strconv.Atoi(str)
	ErrLog(err)
	return re
}

func StrToInt64(str string) int64 {
	re, err := strconv.ParseInt(str, 10, 64)
	ErrLog(err)
	return re
}

func StrToFloat64(str string) (float64, error) {
	re, err := strconv.ParseFloat(str, 64)
	//ErrLog(err)
	return re, err
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FloatToStr(f float64) string {
	//return strconv.FormatFloat(f, 'E', -1, 64)
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func StructToMap(st interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	j, err := jsoniter.Marshal(st)
	if err != nil {
		return nil, err
	}
	err = jsoniter.Unmarshal(j, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func MsqDataToDashSlice(data map[int]map[string]string) godash.DashSlice {
	var nd []map[string]string
	for _, dt := range data {
		nd = append(nd, dt)
	}
	return godash.Flatten(godash.DashSlice{nd})
}

func MsqStructToMap(st interface{}, options ...SMsqStructToMapOptions) map[string]interface{} {
	m, _ := StructToMap(st)
	n := make(map[string]interface{})

	var opt SMsqStructToMapOptions
	if len(options) == 1 {
		opt = options[0]
	}

	for i, v := range m {
		switch v.(type) {
		case string:
			if !StrIsEmpty(v.(string)) {
				n[i] = v
			} else {
				if opt.StrAllowEmpty && godash.Includes(opt.StrAllowEmptyKey, i) {
					n[i] = ""
				}
			}
			break
		case map[string]interface{}:
			//json, _ := jsoniter.Marshal(v)
			//n[i] = string(json)
			n[i], _ = jsoniter.Marshal(v)
			break
		default:
			n[i] = v
		}

		//fmt.Println(b, 22)
	}

	return n
}

func ArrToDashSlice(arr interface{}) godash.DashSlice {
	return godash.Flatten(godash.DashSlice{arr})
}

func ConvertStringEncode(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func ConvertStringUTF8(src string) string {
	return ConvertStringEncode(src, "utf-8", "utf-8")
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if  data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i + 1] >= 0x40 &&
				data[i + 1] <= 0xfe &&
				data[i + 1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i:=0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

func isUtf8(data []byte) bool {
	i := 0
	for i < len(data)  {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num - 1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else  {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func ConvertStringGBKToUTF8(src string) string {
	srcByte := []byte(src)

	if isUtf8(srcByte) {
		return src
	}

	return ConvertStringEncode(src, "gbk", "utf-8")
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//func StrToFloat(str string) float64 {
//	return strconv.ParseFloat(str, 64)
//}

//func Str2TimeDuration(str string) int {
//	re, err := strconv.(str)
//	if err != nil {
//		log.Fatalf("utils convert str2int, fail to parse '%v': %v", str, err)
//	}
//	return re
//}
