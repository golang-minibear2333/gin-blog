package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/golang-minibear2333/gin-blog/pkg/version"

	"github.com/golang-minibear2333/gin-blog/pkg/tracer"

	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/model"
	"github.com/golang-minibear2333/gin-blog/internal/routers"
	"github.com/golang-minibear2333/gin-blog/pkg/logger"
	"github.com/golang-minibear2333/gin-blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	flagRun()
	// 配置初始化，读取到全局model里面
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	// 数据库orm配置初始化
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/golang-minibear2333/gin-blog
func main() {
	global.Logger.Infof("%s: golang-minibear2333/%s", "project", "gin-blog")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
func flagRun() {
	flag.Parse()
	version.CmdParseVersion()
}
func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	// 千万小心全局变量不要被局部变量覆盖，也就是不要用 := 符号
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"gin-blog",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	// 放到全局变量中，供后续中间件或者不同的自定义Span中打点使用
	global.Tracer = jaegerTracer
	return nil
}
