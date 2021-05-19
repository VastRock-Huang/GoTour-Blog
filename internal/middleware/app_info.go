package middleware
//服务信息存储中间件

import "github.com/gin-gonic/gin"

//记录应用信息字段到上下文
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name","GoTour-BlogService")
		c.Set("app_version","1.0.0")
		c.Next()
	}
}
