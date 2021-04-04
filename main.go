package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/internal/routers"
	"github.com/vastrock-huang/gotour-blogservice/pkg/setting"
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
}

func main() {
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

func setupDBEngine() error {
	var err error
	global.DBEngine,err=model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
