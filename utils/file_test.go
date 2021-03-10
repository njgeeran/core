package utils

import (
	"fmt"
	"testing"
)

func TestDeCompressZip(t *testing.T) {
	err := DeCompressZip("C:/Users/Administrator/Desktop/W2Deploy/CUST_CUSTOMER/Desktop.zip","C:/Users/Administrator/Desktop/W2Deploy/CUST_CUSTOMER/Desktop/")
	if err != nil {
		fmt.Println(err.Error())
	}
}
