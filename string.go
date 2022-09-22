package ndash

import (
	"github.com/ginkgoch/godash"
	"math/rand"
	"net/url"
	"strings"
)

func StrIsEmpty(str string) bool {
	//if &str == nil {
	//	return true
	//}
	return len(str) == 0
}

func StrTrim(str string, keys ...string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)

	return str
}

func StrTrimLeft(str string, keys ...string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimLeft(str, "\n")

	return str
}

func StrTrimRight(str string, keys ...string) string {
	str = strings.TrimRight(str, " ")
	str = strings.TrimRight(str, "\n")

	return str
}

func StrTrimBoth(str string, keys ...string) string {
	str = StrTrimLeft(str)
	str = StrTrimRight(str)

	return str
}

func StrSplitLen(str string, length int) []string {
	var strSlice []string
	tmp := str
	i := 0

	var loop func(int)
	loop = func(i int) {
		if len(tmp) > length {
			sl := tmp[0:length]
			strSlice = append(strSlice, sl)
			tmp = tmp[length:]
			loop(i + 1)
		} else {
			strSlice = append(strSlice, tmp)
		}
	}

	loop(i)

	return strSlice
}

func UriEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func UriDecode(str string) string {
	s, err := url.QueryUnescape(str)
	ErrLog(err)
	return s
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(Timezone8().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func StrReverse(str string) string {
	return godash.Join(godash.Reverse(godash.Flatten(godash.DashSlice{godash.Split(str, "")})), "")
}

func StrSort(str string) string {
	return godash.Join(ArrSort(godash.Flatten(godash.DashSlice{godash.Split(str, "")})), "")
}
