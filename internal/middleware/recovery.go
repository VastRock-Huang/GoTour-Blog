package middleware
//自定义错误恢复加发送邮件中间件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/email"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
	"time"
)

//自定义恢复中间件Recovery
func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host: global.EmailSetting.Host,
		Port: global.EmailSetting.Port,
		IsSSL: global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From: global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			//出现异常时发送邮件
			if err := recover();err!=nil {
				//日志记录函数调用信息
				global.Logger.WithCallerFrames().Errorf(c,"panic recover err: %v",err)
				//发送邮件
				if err := defaultMailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发生时间: %v",time.Now()),
					fmt.Sprintf("错误信息: %v",err),
					); err !=nil {	//若发送邮件错误,且此处是由异常引发的
					global.Logger.Panicf(c,"mail.SendMail err: %v",err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()	//终止处理程序
			}
		}()
		c.Next()
	}
}
