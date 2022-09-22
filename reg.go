package ndash

import (
	"fmt"
	"regexp"
)

func RegMatch(regex string, str string) bool {
	reg, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(err)
	}

	return reg.MatchString(str)
}

func RegExec(regex string, str string) [][]string {
	var result [][]string
	reg := regexp.MustCompile(regex)
	if reg == nil {
		fmt.Println("MustCompile err")
	} else {
		result = reg.FindAllStringSubmatch(str, -1)
	}

	return result
}

func RegReplaceAllStringFunc(regex string, str string, iterator func(string) string) string {
	re3, _ := regexp.Compile(regex)
	rep := re3.ReplaceAllStringFunc(str, iterator)
	return rep
}
