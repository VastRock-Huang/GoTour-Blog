package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/service"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
)

// @Summary 校验获取认证信息
// @Produce json
// @Param app_key query string true "app_key" maxlength(20)
// @Param app_secret query string true "app_secret" maxlength(50)
// @Success 200 {object} model.AuthSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 401 {object} errcode.Error "未授权"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	if valid,errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf(c,"app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c)
	if err := svc.CheckAuth(&param); err != nil {
		global.Logger.Errorf(c,"svc.CheckAuth err: %v",err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token,err:=app.GenerateToken(param.AppKey,param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c,"app.GenerateToken err: %v",err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token":token,
	})
}