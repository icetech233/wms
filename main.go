package main

import (
	"fmt"
	"net/http"
	v1 "wms/api/v1"
	"wms/internal"
	"wms/middleware"

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
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middleware.Cors())
	// engine.Use(cors.Default())
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1api := engine.Group("/api/v1")
	{
		v1api.GET("/spu/list", v1.SpuV1.SpuList)
		v1api.POST("/spu/add", v1.SpuV1.SpuAdd)
	}
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
		fmt.Println(err)
	}

	// 初始化gplus
	gplus.Init(gormDb)
}
