package utils

import (
	"fmt"
	"testing"
)

func TestCamelCaseToUnderline(t *testing.T) {
	result := CamelCaseToUnderline("UserName")
	fmt.Println(result)
	result = CamelCaseToUnderline("User_name")
	fmt.Println(result)
}
func TestUnderlinToCamelCase(t *testing.T) {
	result := UnderlinToCamelCase("user_name")
	fmt.Println(result)
}
