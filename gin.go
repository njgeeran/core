package core

import (
	"github.com/gin-gonic/gin"
	"github.com/njgeeran/core/conf"
	"github.com/njgeeran/core/log"
	"github.com/njgeeran/core/middleware"
)

func InitGin() *gin.Engine {
	//禁用默认得日志
	gin.DefaultWriter = &log.DisableGinDefaultLog{}
	//修改默认得错误日志
	gin.DefaultErrorWriter = &log.GinErrLog{}
	if conf.GetConf().SysSetting.Env != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	//开启gin
	r := gin.Default()
	// 跨域
	r.Use(middleware.Cors())
	//捕获异常
	r.Use(gin.Recovery())
	//开启日志
	r.Use(middleware.EnableGinLog())
	return r
}