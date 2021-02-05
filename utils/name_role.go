package utils

import (
	"strings"
)

func FirstNameToUpper(str string) string {
	return strings.ToUpper(string(str[0]))+str[1:]
}
//驼峰命名法转化为下划线
func CamelCaseToUnderline (str string) string {
	strs := FindAllString("[A-Z]{1}[a-z]+",str)
	if len(strs) <= 0 {
		return ""
	}
	result := []string{}
	for _,t := range strs {
		result = append(result, strings.ToLower(t))
	}
	return strings.Join(result,"_")
}
func UnderlinToCamelCase(str string) string {
	strs := strings.Split(str,"_")
	result := ""
	for _,t := range strs {
		result += strings.ToUpper(string(t[0]))+t[1:]
	}
	return result
}