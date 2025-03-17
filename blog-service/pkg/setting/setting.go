package setting

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Setting 提供加载配置文件，监控配置文件
type Setting struct {
	vp *viper.Viper
}

// NewSetting 创建Setting结构体。configs参数是多个配置文件路径
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

// WatchSettingChange 监控配置文件变化，发生改变则重新加载配置文件
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
