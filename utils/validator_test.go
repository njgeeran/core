package utils

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T)  {
	type test3 struct {
		Name		string
	}
	type test2 struct {
		Name		string
		Test3		test3
	}
	type test struct {
		Test2 		test2
		Test2Arr	[]test2
		Name		string
		Phone 		string
	}

	tt := test{
		Name:"测试",
		Phone:"13812345679",
		Test2:test2{
			Name:"测试2",
			Test3:test3{
				Name:"测试3",
			},
		},
		Test2Arr:[]test2{
			{
				Name:"测试2Arr",
			},
			{
				Name:"测试2Arr2",
				Test3:test3{
					Name:"",
				},
			},
		},
	}
	var role = map[string][]string{
		"Name":{NotEmpty()},
		"Phone":{NotEmpty(),VerPhone()},
		"Test2.Name":{NotEmpty()},
		"Test2.Test3.Name":{NotEmpty()},
		"Test2Arr.Name":{NotEmpty()},
		"Test2Arr.Test3.Name":{NotEmpty()},
	}
	if err := Verify(tt,role);err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("验证成功")
}
