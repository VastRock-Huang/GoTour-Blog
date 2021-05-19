package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/service"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/convert"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
)

//文章
type Article struct {}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单个文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error	"内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context)  {
	//获取文章请求的参数结构体
	param := service.GetArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	//新建响应上下文
	response:=app.NewResponse(c)
	//绑定参数并验证
	if valid, errs := app.BindAndValid(c, &param); !valid {
		global.Logger.Errorf(c,"app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//新建服务请求
	svc := service.New(c.Request.Context())
	article,err:= svc.GetArticle(&param)	//获取文章
	if err != nil {
		global.Logger.Errorf(c,"svc.GetArticle err: %v",err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)	//返回文章的响应
}

// @Summary 获取多个文章
// @Produce json
// @Param name query string false "文章名称" minlength(2) maxlength(100)
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc:=service.New(c.Request.Context())
	//设置分页信息
	pager := app.Pager{
		Page: app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	//获取文章列表,返回文章列表,总数
	articles, totalRows, err:=svc.GetArticleList(&param,&pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err: %v",err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	response.ToResponseList(articles,totalRows)
}


// @Summary 创建文章
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param title body string true "文章标题" minlength(2) maxlength(100)
// @Param desc body string false "文章简述" minlength(2) maxlength(255)
// @Param cover_image_url body string true "封面图片地址" maxlength(255)
// @Param content body string true "文章内容" minlength(2) maxlength(4294967295)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context)  {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateArticle(&param); err != nil {
		global.Logger.Errorf(c,"svc.CreateArticle err: %v",err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章ID"
// @Param tag_id body string false "标签ID"
// @Param title body string false "文字标题"	minlength(2) maxlength(100)
// @Param desc body string false "文章简述" minlength(2) maxlength(255)
// @Param cover_image_url body string false "封面图片地址" maxlength(255)
// @Param content body string false "文章内容" minlength(2), maxlength(4294967295)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(2) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context)  {
	param := service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.UpdateArticle(&param); err != nil {
		global.Logger.Errorf(c,"svc.UpdateArticle err: %v",err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context)  {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c,&param); !valid {
		global.Logger.Errorf(c,"app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err:= svc.DeleteArticle(&param); err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v",err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}
