package setting

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp: vp}

	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			log.Println("config changed, reload")
			// 监测到配置文件发生变化，重新加载配置文件
			_ = s.ReloadAllSection()
		})
	}()
}
