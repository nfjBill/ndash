package ndash

import "github.com/ginkgoch/godash"

func ArrJoin(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func ArrMap(arr []string, callback func(v string) string) []string {
	var na []string
	na = make([]string, len(arr))
	for i, v := range arr {
		na[i] = callback(v)
	}
	return na
}

func ArrSort(a godash.DashSlice) godash.DashSlice {
	sort := godash.SortByString(a, func(i interface{}) string {
		return i.(string)
	})

	return sort
}