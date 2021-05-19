package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
	"net/http"
)

//响应
type Response struct {
	Ctx *gin.Context
}

//分页
type Pager struct {
	Page      int `json:"page"`       //页号
	PageSize  int `json:"page_size"`  //页大小
	TotalRows int `json:"total_rows"` //总记录数
}

//新建响应上下文
func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

//设置列表(列出标签,列出文章)响应
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list, //记录列表
		"pager": Pager{ //分页信息
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows, //页内记录数
		},
	})
}

//设置错误响应
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	details := err.Details()
	//若有详情信息则添加到响应中
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
