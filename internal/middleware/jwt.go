package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
)

//JWT鉴权中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			errCode = errcode.Success
		)
		//判断token是否在请求参数或首部字段中
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else if s = c.PostForm("token"); s != ""{
			token = s	//检查是否在请求表单中
		} else {
			token = c.GetHeader("token")
		}
		//未找到token则参数无效
		if token == "" {
			errCode = errcode.InvalidParams
		} else {
			//解析校验token
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:	//JTW时效错误
					errCode = errcode.UnauthorizedTokenTimeout
				default:
					errCode = errcode.UnauthorizedTokenError
				}
			}
		}
		//若认证有错误发送错误响应
		if errCode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(errCode)
			c.Abort()	//终止后续处理函数的执行
			return
		}
		c.Next()
	}
}
