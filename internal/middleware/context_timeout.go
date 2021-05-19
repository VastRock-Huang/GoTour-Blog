package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

//超时控制中间件
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//设置上下文超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		//将新的上下文赋值给请求,超时后会取消请求
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
