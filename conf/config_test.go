package conf

import (
	"fmt"
	"testing"
)

func TestInitConfg(t *testing.T) {
	conf := InitConfg()
	fmt.Println(conf)
}

func TestSettings_GetChildd(t *testing.T) {
	conf := InitConfg()
	key := conf.Setting.GetChildd("test").GetStringd("key","default_value")
	fmt.Println(key)
}