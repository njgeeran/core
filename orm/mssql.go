package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/njgeeran/core/conf"
	"github.com/njgeeran/core/log"
	"net/url"
	"os"
)

func InitMsSql(mssql_s *conf.Settings,orms *Orm) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("初始化mssql失败：%s\n", r)
			return
		}
	}()
	if mssql_s == nil {
		return
	}

	for k,_ := range *mssql_s {
		s := mssql_s.GetChild(k)
		username,password,path := s.GetStringd("username","sa"),s.GetString("password"),s.GetStringd("path","127.0.0.1:1433")
		db_name := s.GetString("db_name")
		max_idle_conns,max_open_conns := s.GetIntd("max-idle-conns",10),s.GetIntd("max-open-conns",10)
		log_mode := s.GetBoold("log_mode",true)
		if db := init_mssql(username,password,path,db_name,max_idle_conns,max_open_conns,log_mode);db != nil{
			orms.Set(k,db)
		}
	}
}

func init_mssql(username,password,path,db_name string,max_idle_conns,max_open_conns int,log_mode bool) *gorm.DB {
	if db, err := gorm.Open("mssql", "sqlserver://"+username+":"+url.QueryEscape(password)+"@"+path+"?database="+db_name); err != nil {
		fmt.Println("MsSQL启动异常:"+err.Error())
		os.Exit(0)
		return nil
	} else {
		db.DB().SetMaxIdleConns(max_idle_conns)
		db.DB().SetMaxOpenConns(max_open_conns)
		db.LogMode(log_mode)
		if log_mode {
			db.SetLogger(&log.GormLogger{})
		}
		return db
	}
}