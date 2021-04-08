package model

import "github.com/vastrock-huang/gotour-blogservice/pkg/app"

//标签信息
type Tag struct {
	*Model
	Name string `json:"name"`	//标签名
	State uint8 `json:"state"`	//标签状态
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

//获取标签表名
func (t Tag) TableName() string {
	return "blog_tag"
}