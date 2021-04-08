//国际化处理中间件
//将validator验证器的错误信息进行翻译
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(),zh.New(),zh_Hant_TW.New())	//创建国际化翻译中间件
		locale := c.GetHeader("locale")	//获取请求首部语言参数
		trans,_ := uni.GetTranslator(locale)	//返回给定语言的翻译器
		//获取验证器的设置和缓存
		v,ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			//将翻译器注册到验证器中
			switch locale {
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(v,trans)
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(v,trans)
			default:
				_ = zhTranslations.RegisterDefaultTranslations(v,trans)
			}
			//将翻译器存储到gin上下文中
			c.Set("trans",trans)
		}
		c.Next()	//执行下个中间件
	}
}
