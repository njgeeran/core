package utils

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T)  {
	type test struct {
		Name		string
		Phone 		string
	}
	tt := test{
		Name:"1",
		Phone:"13812345679",
	}
	var role = map[string][]string{
		"Name":{NotEmpty()},
		"Phone":{NotEmpty(),VerPhone()},
	}
	if err := Verify(tt,role);err != nil{
		fmt.Println(err)
	}
}
