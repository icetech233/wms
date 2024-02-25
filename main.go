package main

import (
	"fmt"
	"net/http"
	v1 "wms/api/v1"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/spf13/viper/remote"
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
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1api := engine.Group("/api/v1")
	{
		v1api.GET("/spu/list", v1.SpuV1.SpuList)

	}
	engine.Run(":8081")
}
func initViper() {
	v := viper.New()
	v.AddRemoteProvider("consul", "127.0.0.1:8500", "wms")
	v.SetConfigType("yaml")
	err := v.ReadRemoteConfig()
	if err != nil {
		panic("viper err" + err.Error())
	}
	// 读取配置
	mysql_dsn = v.GetString("sql.dsn")
}

// func spuInsert() {
// 	// spu := Spu{SpuCode: "IIG1000E", SpuDesc: "IIG1000E系列"}
// 	spu := Spu{SpuCode: "IIG3000", SpuDesc: "IIG3000系列"}
// 	resultDb := gplus.Insert[Spu](&spu)
// 	spew.Dump(resultDb.Error, resultDb.RowsAffected)
// 	fmt.Println("spu id:", spu.SpuID)
// }

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
