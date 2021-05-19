package dao

import (
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
)

//给定相关参数创建文章
func (d *Dao) CreateArticle(title, desc, content, coverImgUrl string,
	state uint8, createdBy string) (*model.Article,error) {
	article:=model.Article{
		Title: title,
		Desc: desc,
		Content: content,
		CoverImageUrl: coverImgUrl,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return article.Create(d.engine)
}

//给定参数更新文章
func (d *Dao) UpdateArticle(id uint32,title, desc, content, coverImgUrl string,
	state uint8, modifiedBy string) error {
	article:=model.Article{
		Model: &model.Model{
			ID: id,
		},
	}
	//指定具体修改的字段
	//使用values映射传递具体要修改的字段
	values:=map[string]interface{}{
		"modified_by": modifiedBy,
		"state": state,
	}
	if title != "" {
		values["title"]=title
	}
	if coverImgUrl != "" {
		values["cover_image_url"]=coverImgUrl
	}
	if desc != "" {
		values["desc"]=desc
	}
	if content != "" {
		values["content"]=content
	}
	return article.Update(d.engine,values)
}

//给定文章id和状态获取相应的文章
func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article:=model.Article{
		State: state,
		Model:&model.Model{
			ID: id,
		},
	}
	return article.Get(d.engine)
}

//给定文章id删除文章
func (d *Dao) DeleteArticle(id uint32) error {
	article:=model.Article{
		Model:&model.Model{
			ID: id,
		},
	}
	return article.Delete(d.engine)
}

//给定标签ID和文章状态统计该标签的文章数
func (d *Dao) CountArticleListByTagID(tagID uint32, state uint8) (int, error) {
	article := model.Article{
		State: state,
	}
	return article.CountByTagID(d.engine,tagID)
}

//给定标签ID,文章状态,页号和页大小获得该标签的文章行信息
func (d *Dao) GetArticleListByTagID(tagID uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{
		State: state,
	}
	return article.ListByTagID(d.engine,tagID,app.GetPageOffset(page,pageSize),pageSize)
}