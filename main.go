package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/internal/routers"
	"github.com/vastrock-huang/gotour-blogservice/pkg/logger"
	"github.com/vastrock-huang/gotour-blogservice/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// 初始化函数
func init() {
	err := setupSetting()	//初始化配置
	if err!=nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description GoTour-BlogService
// @termsOfService https://github.com/vastrock-huanng/gotour-blogservice
func main() {
	global.Logger.Infof("%s: gotour-blog/%s","hhh","blog_service")

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,	//处理接口调用ServeHTTP
		ReadTimeout: global.ServerSetting.ReadTimeout,	//请求读取操作的最大时间
		WriteTimeout: global.ServerSetting.WriteTimeout,	//恢复写入操作的最大时间
		MaxHeaderBytes: 1<<20,	//请求首部最大字节数
	}
	_ = s.ListenAndServe()	//监听并服务
	//router.Run()
}

//初始化配置
func setupSetting() error {
	//初始化配置对象,内置Viper解析配置文件
	set,err:=setting.NewSetting()
	if err!=nil{
		return err
	}
	//读取配置到全局配置变量
	err = set.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout*=time.Second
	global.ServerSetting.WriteTimeout*=time.Second
	err = set.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

//初始化数据库
func setupDBEngine() error {
	var err error
	global.DBEngine,err=model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

//初始化日志
func setupLogger() error {
	//日志文件路径
	fileName := global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+
		global.AppSetting.LogFileExt
	global.Logger=logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,	//日志文件路径
		MaxSize: 600,		//日志文件最大空间(MB)
		MaxAge: 10,			//日志文件最大生存周期(天)
		LocalTime: true,
	},"",log.LstdFlags).	//标准日志参数
		WithCaller(2)	//跳过2个调用栈,即栈信息为setupLogger的调用者init函数
	return nil
}