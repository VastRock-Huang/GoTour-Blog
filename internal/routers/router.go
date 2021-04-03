package routers

import "github.com/gin-gonic/gin"

//新建路由
func NewRouter() *gin.Engine {
	r:=gin.New()
	r.Use(gin.Recovery(),gin.Logger())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/tags")		//新增标签
		apiV1.DELETE("/tags/:id")//删除标签
		apiV1.PUT("/tags/:id")	//更新标签
		apiV1.GET("/tags")		//获取标签列表

		apiV1.POST("/articles")		//新增文章
		apiV1.DELETE("/articles/:id")//删除文章
		apiV1.PUT("/articles/:id")	//更新文章
		apiV1.GET("/articles/:id")	//获取文章
		apiV1.GET("/articles")		//获取文章列表
	}
	return r
}

