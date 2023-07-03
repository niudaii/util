package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func New(filename string, val interface{}) (err error) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("info %v config changed\n", e.Name)
		if err = v.Unmarshal(&val); err != nil {
			return
		}
		log.Printf("info %v 解析成功\n", filename)
	})
	v.WatchConfig()
	err = v.Unmarshal(&val)
	return
}
