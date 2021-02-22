package utils

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	str := RandomString(10)
	fmt.Println(str)
}
