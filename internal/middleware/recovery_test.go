package middleware

import (
	"errors"
	"fmt"
	"github.com/vastrock-huang/gotour-blogservice/pkg/email"
	"testing"
	"time"
)

func TestRecovery(t *testing.T) {
	var defaultMailer = email.NewEmail(&email.SMTPInfo{
		Host:     "smtp.qq.com",
		Port:     465,
		IsSSL:    true,
		UserName: "xxx@qq.com",
		Password: "xxx",
		From:     "xxx@qq.com",
	})
	if err := errors.New("test err"); err != nil {
		//global.Logger.WithCallerFrames().Errorf("panic recover err: %v",err)
		if err := defaultMailer.SendEmail(
			[]string{"xxx@qq.com"},
			fmt.Sprintf("异常抛出,发生时间: %v", time.Now()),
			fmt.Sprintf("错误信息: %v", err),
		); err != nil {
			fmt.Printf("mail.SendMail err: %v", err)
		}
	}
}
