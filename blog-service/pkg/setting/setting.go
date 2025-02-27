package setting

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config changed")
	})
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
