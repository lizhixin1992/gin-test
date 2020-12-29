package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobalConf = New()

func New() *viper.Viper {
	v := viper.New()
	////toml格式
	//v.SetConfigName("GlobalConf")
	//v.SetConfigType("toml")
	//v.AddConfigPath("./conf/")

	//yml格式
	v.SetConfigName("config_test")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf/")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}
