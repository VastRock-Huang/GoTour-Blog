package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
	"github.com/vastrock-huang/gotour-blogservice/pkg/limiter"
)

//限流中间件
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)		//获取对应令牌桶关键字,即去参URI
		//获取关键字对应的令牌桶
		//此处以路径来做关键字,即若该路径有限流器,则进行限流操作,否则直接跳过
		if bucket,ok:=l.GetBucket(key);ok{
			//取走1个令牌
			count := bucket.TakeAvailable(1)
			//无令牌返回错误
			if count == 0 {
				response:=app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
