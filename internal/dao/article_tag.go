package dao

import "github.com/vastrock-huang/gotour-blogservice/internal/model"

//给定文章ID获取一条文章标签记录
func (d *Dao) GetArticleTagByArticleID(articleID uint32) (model.ArticleTag,error) {
	articleTag:=model.ArticleTag{
		ArticleID: articleID,
	}
	return articleTag.GetByArticleID(d.engine)
}

//跟定标签ID获取文章标签记录列表
func (d *Dao) GetArticleTagListByTagID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag:=model.ArticleTag{
		TagID: tagID,
	}
	return articleTag.ListByTagID(d.engine)
}

//给定一组文章ID获取文章标签记录列表
func (d *Dao) GetArticleTagListByArticleIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag:=model.ArticleTag{}
	return articleTag.ListByArticleIDs(d.engine,articleIDs)
}

//给定文章ID,标签ID以及创建者创建一条文章标签记录
func (d *Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag:=model.ArticleTag{
		ArticleID: articleID,
		TagID: tagID,
		Model:&model.Model{
			CreatedBy: createdBy,
		},
	}
	return articleTag.Create(d.engine)
}

//给定文章ID,标签ID以及修改者修改一套文章标签记录
func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{}
	values:=map[string]interface{}{
		"article_id":articleID,
		"tag_id": tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine,values)
}

//给定文章ID删除一条文章标签记录
func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag:=model.ArticleTag{
		ArticleID: articleID,
	}
	return articleTag.DeleteOne(d.engine)
}