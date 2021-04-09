package model

import "github.com/vastrock-huang/gotour-blogservice/pkg/app"

//文章信息
type Article struct {
	*Model
	Title         string `json:"title"`           //文章标题
	Desc          string `json:"desc"`            //文章简述
	CoverImageUrl string `json:"cover_image_url"` //文章封面url
	Content       string `json:"content"`         //文章内容
	State         uint8  `json:"state"`           //文章状态
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

//获取文章表名
func (a Article) TableName() string {
	return "blog_article"
}
