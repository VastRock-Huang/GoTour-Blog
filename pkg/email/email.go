package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

//SMTP信息
type SMTPInfo struct {
	Host string		//主机
	Port int		//端口
	IsSSL bool		//是否支持SSL
	UserName string	//用户名
	Password string	//密码
	From string
}

//邮件
type Email struct {
	*SMTPInfo
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{
		SMTPInfo:info,
	}
}

//发送邮件
func (e *Email) SendEmail(to []string, subject, body string) error {
	m:=gomail.NewMessage()	//新建邮件信息
	//设置邮件字段
	m.SetHeader("From",e.From)	//发送者
	m.SetHeader("To",to...)	//接受者
	m.SetHeader("Subject",subject)	//主题
	m.SetBody("text/html",body)	//内容
	//创建一个新SMTP对话
	dialer := gomail.NewDialer(e.Host, e.Port,e.UserName,e.Password)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: e.IsSSL,
	}
	return dialer.DialAndSend(m)	//发送邮件信息
}