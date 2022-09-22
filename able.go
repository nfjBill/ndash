package ndash

import (
	"github.com/ginkgoch/godash"
	"strings"
)

func UrlQueryStr(url string) string {
	var str string
	urlArr := strings.Split(url, "?")

	if len(urlArr) > 1 {
		str = urlArr[1]
	}

	if !StrIsEmpty(str) {
		str = strings.Split(str, "#")[0]
	}

	return str
}

func UrlQueryIterateToStr(queryStr string, callback func(v string, key string, source string) string) string {
	if !StrIsEmpty(queryStr) {
		arr := strings.Split(queryStr, "&")

		nArr := godash.Map(godash.Flatten(godash.DashSlice{arr}), func(i interface{}) interface{} {
			var r string
			v := i.(string)
			r = v
			j := strings.Index(v, "=")
			eArr := []string{v}
			if j != -1 {
				eArr = []string{v[:j], v[j+1:]}
			}

			if len(eArr) > 1 {
				c := callback(eArr[1], eArr[0], v)
				r = eArr[0] + "=" + c
			} else {
				c := callback(eArr[0], eArr[0], v)
				r = c
			}
			return r
		})

		queryStr = godash.Join(nArr, "&")
	}

	return queryStr
}

func UrlQueryToMap(queryStr string) map[string][]string {
	var query = make(map[string][]string)
	if !StrIsEmpty(queryStr) {
		//if !strings.HasPrefix("?", queryStr) {
		//	queryStr = "?" + queryStr
		//}
		UrlQueryIterateToStr(queryStr, func(v string, key string, source string) string {
			query[key] = append(query[key], v)
			return ""
		})
	}
	return query
}

func UrlQueryMapToString(queryMap map[string][]string) string {
	var query string
	for key, arrVal := range queryMap {
		for _, v := range arrVal {
			if !StrIsEmpty(query) && strings.HasSuffix(query, "&") {
				query += "&"
			}

			query += key + "=" + UriEncode(v)
		}
	}

	return query
}

func UrlQuerySingleValue(key string, queryMap map[string][]string) string {
	arr := queryMap[key]

	if len(arr) > 0 {
		return arr[0]
	}

	return ""
}

func FormDataFieldMap(field map[string][]string, callback func(val string, key string) string) map[string][]string {
	var fieldData = make(map[string][]string)

	for n, m := range field {
		for _, f := range m {
			fieldData[n] = append(fieldData[n], callback(f, n))
		}
	}

	return fieldData
}
