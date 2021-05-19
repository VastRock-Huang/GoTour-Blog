package model

import "github.com/jinzhu/gorm"

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

//根据文章ID获取一条文章标签记录
func (a ArticleTag) GetByArticleID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	// SELECT * FROM article_tag
	// WHERE article_id=a.ArticleID AND is_del=0
	// ORDER BY id LIMIT 1;
	if err:=db.Where("article_id=? AND is_del=?",a.ArticleID,0).
		First(&articleTag).Error; err!=nil && err!=gorm.ErrRecordNotFound {
			return articleTag,err
	}
	return articleTag,nil
}

//根据标签ID获取文章标签记录列表
func (a ArticleTag) ListByTagID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	// SELECT * FROM article_tag WHERE tag_id=a.TagID AND is_del=0;
	if err:= db.Where("tag_id=? AND is_del=?",a.TagID,0).
		Find(&articleTags).Error; err!=nil {
			return nil, err
	}
	return articleTags,nil
}

//根据一组文章ID获取文章标签记录列表
func (a ArticleTag) ListByArticleIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	// SELECT * FROM article_tag WHERE article_id IN articleIDs AND is_del=0;
	if err:=db.Where("article_id IN (?) AND is_del=?",articleIDs,0).
		Find(&articleTags).Error; err!=nil && err!=gorm.ErrRecordNotFound {
			return nil, err
	}
	return articleTags,nil
}

//创建一条文章标签记录
func (a ArticleTag) Create(db *gorm.DB) error {
	// INSERT INTO article_tag VALUES(...);
	if err:=db.Create(&a).Error;err!=nil{
		return err
	}
	return nil
}

//根据文章ID更新一条文章标签记录
func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	// UPDATE article_tag SET ...
	// WHERE article_id=a.ArticleID AND is_del=0 LIMIT 1;
	if err:=db.Model(&ArticleTag{}).Where("article_id=? AND is_del=?",a.ArticleID,0).
		Limit(1).Updates(values).Error; err!=nil {
			return err
	}
	return nil
}

//删除文章标签记录
func (a ArticleTag) Delete(db *gorm.DB) error {
	// DELETE FROM article_tag WHERE id=a.Model.ID AND is_del=0;
	if err:=db.Where("id=? AND is_del=?",a.Model.ID,0).
		Delete(&a).Error;err!=nil{
			return err
	}
	return nil
}

//根据文章ID删除一条文章标签记录
func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	// DELETE FROM article_tag WHERE article_id=a.ArticleID AND is_del=0 LIMIT 1;
	if err:=db.Where("article_id=? AND is_del=?",a.ArticleID,0).
		Delete(&ArticleTag{}).Limit(1).Error;err!=nil{
			return err
	}
	return nil
}