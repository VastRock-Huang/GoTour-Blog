package model

//文章和标签关联信息
type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id"`//文章ID
	TagID uint32 `json:"tag_id"`		//标签ID
}

//获取文章标签关联表名
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}