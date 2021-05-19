package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/vastrock-huang/gotour-blogservice/docs"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/middleware"
	"github.com/vastrock-huang/gotour-blogservice/internal/routers/api"
	"github.com/vastrock-huang/gotour-blogservice/internal/routers/api/v1"
	"github.com/vastrock-huang/gotour-blogservice/pkg/limiter"
	"net/http"
	"time"
)

//路由限流器
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",	//对该路径进行限流
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

//新建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	// 日志记录中间件和错误恢复中间件
	//if global.ServerSetting.RunMode == "debug" {
	//	r.Use(gin.Logger())
	//	r.Use(gin.Recovery())
	//}else{
	//	r.Use(middleware.AccessLog())
	//	r.Use(middleware.Recovery())
	//}
	r.Use(middleware.AccessLog(),middleware.Recovery())

	r.Use(middleware.Tracing())	//链路追踪中间件
	r.Use(middleware.RateLimiter(methodLimiters))	//限流器中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))	//超时控制中间件

	r.Use(middleware.Translations())	//注册验证器的多语言中间件
	//接口文档路径
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))
	//上传文件
	upload := api.NewUpload()
	r.POST("/upload/file",upload.UploadFile)
	//静态资源访问
	r.StaticFS("/static",http.Dir(global.AppSetting.UploadSavePath))
	//授权认证
	r.GET("/auth",api.GetAuth)

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group("/api/v1")
	//apiV1.Use(middleware.JWT())		//鉴权中间件
	{
		apiV1.POST("/tags", tag.Create)       //新增标签
		apiV1.DELETE("/tags/:id", tag.Delete) //删除标签
		apiV1.PUT("/tags/:id", tag.Update)    //更新标签
		apiV1.PATCH("/tags/:id/state",tag.Update)		//更新标签一部分
		apiV1.GET("/tags", tag.List)           //获取标签列表

		apiV1.POST("/articles", article.Create)       //新增文章
		apiV1.DELETE("/articles/:id", article.Delete) //删除文章
		apiV1.PUT("/articles/:id", article.Update)    //更新文章
		apiV1.PATCH("/articles/:id/state",article.Update)
		apiV1.GET("/articles/:id", article.Get)       //获取文章
		apiV1.GET("/articles", article.List)          //获取文章列表
	}
	return r
}
