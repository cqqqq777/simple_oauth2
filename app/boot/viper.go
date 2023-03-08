package boot

import (
	"fmt"
	g "github.com/cqqqq777/simple_oauth2/app/global"
	"github.com/spf13/viper"
)

const (
	configPath = "app/config/config.yaml"
)

func ViperInit() {
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed err:%v", err))
	}
	if err := v.Unmarshal(&g.Config); err != nil {
		panic(fmt.Errorf("unmarshal config failed err:%v", err))
	}
}
