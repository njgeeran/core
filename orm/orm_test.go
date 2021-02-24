package orm

import (
	"fmt"
	"github.com/njgeeran/core/conf"
	"testing"
)

func TestInitOrm(t *testing.T) {
	config := conf.InitConfg()
	orms := InitOrm(config)
	fmt.Println(orms.dbs)
}