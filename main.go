package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/internal/routers"
	"github.com/vastrock-huang/gotour-blogservice/pkg/logger"
	"github.com/vastrock-huang/gotour-blogservice/pkg/setting"
	"github.com/vastrock-huang/gotour-blogservice/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	//命令行参数对应的变量
	port string		//服务端口号
	runMode string	//运行模式
	config string	//配置文件路径

	//编译信息
	isVersion bool		//是否显示版本信息
	buildTime string	//编译时间
	buildVersion string	//编译版本
	gitCommitID string	//Git Commit ID
)

// 初始化函数
func init() {
	var err error
	err=setupFlag()	//首先进行命令行解析
	if err!=nil{
		log.Fatalf("init.setupFlag err: %v",err)
	}
	err = setupSetting()	//初始化配置
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
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v",err)
	}
}

////检查热更新情况
//func checkHotUpdate()  {
//	go func(){
//		for {
//			<-setting.Ch
//			log.Printf("globa.ServerSeting: %+v",global.ServerSetting)
//			//log.Printf("global.AppSetting: %+v",global.AppSetting)
//		}
//	}()
//}

// @title 博客系统
// @version 1.0
// @description GoTour-BlogService
// @termsOfService https://github.com/vastrock-huanng/gotour-blogservice
func main() {
	//global.Logger.Infof("%s: gotour-blog/%s","hhh","blog_service")
	//checkHotUpdate()

	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n",buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}


	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,	//处理接口调用ServeHTTP
		ReadTimeout: global.ServerSetting.ReadTimeout,	//请求读取操作的最大时间
		WriteTimeout: global.ServerSetting.WriteTimeout,	//恢复写入操作的最大时间
		MaxHeaderBytes: 1<<20,	//请求首部最大字节数
	}

	//开启新的goroutine进行监听并服务
	go func() {
		//若启动服务错误且不为服务被关闭,则报错终止程序
		if err:=s.ListenAndServe(); err!=nil && err!=http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v",err)
		}
	}()
	//等待中断信号
	quit := make(chan os.Signal)	//程序中断的管道
	//收到SIGINT(中断信号Ctrl+C)或SIGTERM信号时将信号发送给quit管道
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	//等待信号,阻塞到信号到来
	<-quit
	log.Println("Shutting down server...")

	//设置带超时处理的上下文
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()	//确保在上下文超时关闭前释放资源
	//优雅地关闭服务(最长等待5秒后会关闭服务)
	if err := s.Shutdown(ctx); err!= nil {
		log.Fatalf("Server forced to shutdown: %v",err)
	}
	log.Println("Server exiting")
}

//初始化配置
func setupSetting() error {
	//初始化配置对象,内置Viper解析配置文件
	//传参为配置文件路径(多个,有优先级)
	set,err:=setting.NewSetting(strings.Split(config,",")...)
	if err!=nil{
		return err
	}
	//读取配置到全局配置变量
	err = set.ReadSection("Server",&global.ServerSetting)
	if err != nil{
		return err
	}
	//global.ServerSetting.ReadTimeout*=time.Second
	//global.ServerSetting.WriteTimeout*=time.Second
	err = set.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	//global.AppSetting.DefaultContextTimeout*=time.Second
	err = set.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("JWT",&global.JWTSetting)
	if err != nil {
		return err
	}
	//global.JWTSetting.Expire*=time.Second
	err = set.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	//若程序命令行参数中有端口或运行模式,则以命令行中参数为准(覆盖配置文件中的配置)
	if port != ""{
		global.ServerSetting.HttpPort=port
	}
	if runMode!=""{
		global.ServerSetting.RunMode=runMode
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
	},"",log.LstdFlags)	//标准日志参数
	// 原代码在此处调用WithCaller函数
	//致使日志中的callers字段总指向main.init函数,而并非log真正调用的位置,因此移到了
	// WithCaller(1)	//跳过2个调用栈,即栈信息为setupLogger的调用者init函数
	return nil
}

//初始化Jaeger链路追踪器
func setupTracer() error {
	//创建Jaeger链路追踪器
	jaegerTracer,_,err:=tracer.NewJaegerTracer(
		"blog-service",	//服务名称
		"127.0.0.1:6831",	//代理主机及端口
		)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer	//绑定到全局
	return nil
}

//程序命令行参数解析
func setupFlag() error {
	flag.BoolVar(&isVersion,"version",false,"编译信息")
	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式")
	flag.StringVar(&config,"config","configs/","指定要使用的配置文件路径")
	//PS: 在powershell中config路径的值要用""括起
	flag.Parse()
	return nil
}