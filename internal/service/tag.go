package service

type CountTagRequest struct {
	Name string `form:"name" binding:"max=100"`	//form为字段名,binding为入参教校验规则
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ListTagRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"name" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"min=2,max=100"`
	State uint8 `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
