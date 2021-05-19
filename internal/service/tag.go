package service
//服务层请求处理
//位于DAO层之上,进行参数验证后将参数传递给DAO

import (
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
)

//请求参数结构体

//统计标签数请求参数
type CountTagRequest struct {
	Name string `form:"name" binding:"max=100"`	//form为字段名,binding为入参教校验规则
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//列出标签请求参数
type TagListRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//创建标签请求参数
type CreateTagRequest struct {
	Name string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//更新标签请求参数
type UpdateTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"max=100"`		//此处不能带min=2限制,否则必须有name参数
	State uint8 `form:"state" binding:"oneof=0 1"`	//此处也无需带default参数
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

//删除标签请求参数
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//统计标签数量
func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name,param.State)
}

//获取标签列表
func (svc *Service) GetTagList(param *TagListRequest,pager *app.Pager) ([]*model.Tag,error) {
	return svc.dao.GetTagList(param.Name,param.State,pager.Page,pager.PageSize)
}

//创建标签
func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name,param.State,param.CreatedBy)
}

//更新标签
func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID,param.Name,param.State,param.ModifiedBy)
}

//删除标签
func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}