package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/spf13/viper/remote"
)

var err error
var gormDb *gorm.DB

func main() {

	initViper()
	return
	initDb()

	// 	viper.SetConfigType("json") // Need to explicitly set this to json
	// err := viper.ReadRemoteConfig()
	// fmt.Println(viper.Get("port")) // 8080
	// fmt.Println(viper.Get("hostname")) // myhostname.com
}

func initViper() {
	v := viper.New()
	v.AddRemoteProvider("consul", "hw.acgzj.cn:8500", "wms")
	v.SetConfigType("yaml")
	err := v.ReadRemoteConfig()

	fmt.Println("viper err", err)
	s1 := v.GetString("dns")

	fmt.Println("s1,", s1)
}

func main2() {

	t1 := time.Now()
	tof := reflect.TypeOf(t1)
	vof := reflect.ValueOf(t1)
	tof = vof.Type()
	fmt.Println(tof.Name(), tof.Kind())
	tof.Elem()

	t2, _ := vof.Interface().(time.Time)
	fmt.Println(t2)
	return

	spus, resultDb := gplus.SelectList[Spu](nil)
	spew.Dump(resultDb.Error, resultDb.RowsAffected)
	for _, spu := range spus {
		b, _ := json.Marshal(spu)
		fmt.Println("spu:", string(b))
	}
}

func spuInsert() {
	// spu := Spu{SpuCode: "IIG1000E", SpuDesc: "IIG1000E系列"}
	spu := Spu{SpuCode: "IIG3000", SpuDesc: "IIG3000系列"}
	resultDb := gplus.Insert[Spu](&spu)
	spew.Dump(resultDb.Error, resultDb.RowsAffected)
	fmt.Println("spu id:", spu.SpuID)
}

func initDb() {
	dsn := "root:271823981@tcp(hw.acgzj.cn:3306)/wms?charset=utf8mb4&parseTime=True&loc=Local"
	gormDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
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
