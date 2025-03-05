package main

import (
	"log"
	"net/http"
	"time"

	"github.com/wushengyouya/blog-service/global"
	"github.com/wushengyouya/blog-service/internal/model"
	"github.com/wushengyouya/blog-service/internal/routers"
	"github.com/wushengyouya/blog-service/pkg/logger"
	"github.com/wushengyouya/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setUpDBEngine err: %v", err)
	}

}

// @title 博客系统
// @version 1.0
// @description Go实现的一个简单的博客项目,参考Go 语言编程之旅
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @termOfService https://github.com/go-programming-tour-book
func main() {
	engine := routers.NewRouters()
	// engine.Run()
	server := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        engine,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	server.ListenAndServe()
}

// 加载配置文件
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// 加载server配置
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	// 加载app配置
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	// 加载数据库配置
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	// 加载jwt配置
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	// 设置超时
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

// 初始化数据库引擎
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
