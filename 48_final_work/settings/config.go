package settings

import (
	"go_web_demo/48_final_work/models"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ViperConfig = new(models.Config)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./settings")
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			return
		}
	}
	if err = viper.Unmarshal(ViperConfig); err != nil {
		return
	}
	//实时监控配置文件的变化
	viper.WatchConfig()
	//当配置变化之后调用的一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件发生变更之后会调用的回调函数
		//记录日志，viper监控的文件出现了变化
		if err = viper.Unmarshal(ViperConfig); err != nil {
			return
		}
	})
	return
}
