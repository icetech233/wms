package main

import (
	"wms/api/v1/router"
	"wms/internal"
	"wms/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysql_dsn string
var err error
var gormDb *gorm.DB

func main() {
	initViper()
	initDb()
	ginStart()
}

func ginStart() {
	// engine.Use(Logger(), Recovery())
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.InitRouterV1(engine)
	engine.Run(":8081")
}
func initViper() {
	v := internal.InitViper()
	// 读取配置
	mysql_dsn = v.GetString("sql.dsn")
}

func initDb() {
	gormDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysql_dsn,
		DisableDatetimePrecision:  true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		util.Log(err)
	}

	// 初始化gplus
	gplus.Init(gormDb)
}
