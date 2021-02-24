package orm

import (
	"fmt"
	"github.com/njgeeran/core/conf"
	"github.com/njgeeran/core/log"
)

//orm 依赖conf和log

func init()  {
	InitOrm(conf.GetConf())
}

func InitOrm(conf *conf.Config) *Orm {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("初始化db失败：%s\n", r)
			return
		}
	}()

	db_s,ch := conf.Setting.GetChildd_c("db")
	orm = NewOrm()
	init_orm(db_s,orm)
	go func() {
		for {
			select {
			case <-ch:
				init_orm(conf.Setting.GetChildd("db"),orm)
				log.GetLoger().Info("监听到数据库配置修改,重新初始化数据库")
			}
		}
	}()
	return orm
}

func init_orm(db_s *conf.Settings,orms *Orm)  {
	for k,_ := range *db_s {
		switch k {
		case "mysql":
			v := db_s.GetChildd("mysql")
			InitMySql(v,orms)
			break
		case "mssql":
			v := db_s.GetChildd("mssql")
			InitMsSql(v,orms)
			break
		}
	}
}