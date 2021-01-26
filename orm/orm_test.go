package orm

import (
	"fmt"
	"github.com/whileW/lowcode-core/conf"
	"testing"
)

func TestInitOrm(t *testing.T) {
	config := conf.InitConfg()
	orms := InitOrm(config)
	fmt.Println(orms.dbs)
}