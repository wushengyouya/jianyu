package global

import (
	"os"
	"path/filepath"
	"sync"
)

var RootDir string
var once = new(sync.Once)

func init() {
	once.Do(func() {
		initConfig()
		inferRootDir()
	})
}

// inferRootDir 推断出项目根目录
func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		// 这里要确保根目录下存在template目录
		if exists(d + "/template") {
			return d
		}
		return infer(filepath.Dir(d))
	}
	RootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
