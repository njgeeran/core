package utils

import (
	"fmt"
	"testing"
)

func TestFindAllString(t *testing.T) {
	result := FindAllString("[A-Z]{1}[a-z]+","UserAge")
	fmt.Println(result)
}
func TestFindString(t *testing.T) {
	result := FindString("[A-Z]{1}[a-z]+","UserAge")
	fmt.Println(result)
}
func TestMatchString(t *testing.T) {
	result := MatchString("[A-Z]{1}[a-z]+","UserAge")
	fmt.Println(result)
}
func TestMatchPhone(t *testing.T) {
	result := MatchPhone("13812345679")
	fmt.Println(result)
}
func TestMatchEmail(t *testing.T) {
	result := MatchEmail("wuwk@geeran.cn")
	fmt.Println(result)
}