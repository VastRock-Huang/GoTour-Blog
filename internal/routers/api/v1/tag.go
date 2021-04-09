package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/service"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/convert"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
)


type Tag struct {

}


func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context)  {
	/* 参数绑定与参数验证部分 */
	param := service.TagListRequest{}	//列出标签请求的参数
	response := app.NewResponse(c)	//响应上下文
	//绑定参数并验证
	if valid, errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		//创建带所有验证错误详情的错误响应
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)	//设置错误响应
		return
	}

	svc := service.New(c.Request.Context())	//新建请求服务
	/* 获得分页信息 */
	//分页信息
	pager := app.Pager{
		Page: app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	//统计总标签数
	totalRows,err:=svc.CountTag(&service.CountTagRequest{
		Name: param.Name,
		State: param.State,
	})
	if err!= nil {
		global.Logger.Errorf("svc.CountTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	/* 获取标签列表 */
	tags, err := svc.GetTagList(&param,&pager)	//获取标签列表
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v",err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags,totalRows)		//设置标签列表的响应
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context)  {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	if valid,errs :=app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	//创建标签
	if err := svc.CreateTag(&param); err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context)  {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),	//从url参数中获取值
	}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	if err:=svc.UpdateTag(&param); err!= nil {
		global.Logger.Errorf("svc.UpdateTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	if valid, errs:= app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp:=errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c)
	if err:=svc.DeleteTag(&param);err!= nil {
		global.Logger.Errorf("svc.DeleteTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
}

