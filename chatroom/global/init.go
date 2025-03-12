package global

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

var RootDir string
var once = new(sync.Once)

func init() {
	once.Do(func() {
		InferRootDir()
		initConfig()
	})
}

// inferRootDir 推断出项目根目录
func InferRootDir() {
	cwd, err := os.Getwd()
	log.Println(cwd)
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
	log.Println(RootDir)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
