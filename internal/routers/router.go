package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/internal/routers/api"
)

//新建路由
func NewRouter() *gin.Engine {
	r:=gin.New()
	r.Use(gin.Recovery(),gin.Logger())

	article := api.NewArticle()
	tag := api.NewTag()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)		//新增标签
		apiV1.DELETE("/tags/:id", tag.Delete)//删除标签
		apiV1.PUT("/tags/:id",tag.Update)	//更新标签
		apiV1.GET("/tags",tag.Get)		//获取标签列表

		apiV1.POST("/articles",article.Create)		//新增文章
		apiV1.DELETE("/articles/:id",article.Delete)//删除文章
		apiV1.PUT("/articles/:id", article.Update)	//更新文章
		apiV1.GET("/articles/:id", article.Get)	//获取文章
		apiV1.GET("/articles",article.List)		//获取文章列表
	}
	return r
}

