package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobalConf = New()

func New() *viper.Viper {
	v := viper.New()
	v.SetConfigName("GlobalConf")
	v.SetConfigType("toml")
	v.AddConfigPath("./conf/")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}
