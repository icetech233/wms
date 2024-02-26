package internal

import (
	"bytes"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	consulVal := readConsul()
	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(bytes.NewReader(consulVal))
	if err != nil {
		panic("viper err" + err.Error())
	}
	return v
}
