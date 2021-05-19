package model

import (
	"github.com/jinzhu/gorm"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
)

//文章信息
type Article struct {
	*Model
	Title         string `json:"title"`           //文章标题
	Desc          string `json:"desc"`            //文章简述
	CoverImageUrl string `json:"cover_image_url"` //文章封面url
	Content       string `json:"content"`         //文章内容
	State         uint8  `json:"state"`           //文章状态
}

//Swagger文档文章信息
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

//获取文章表名
func (a Article) TableName() string {
	return "blog_article"
}

//创建文章
func (a Article) Create(db *gorm.DB) (*Article, error) {
	//INSERT INTO article VALUES(...);
	if err := db.Create(&a).Error; err!=nil{
		return nil,err
	}
	return &a,nil
}

//更新文章
func (a Article) Update(db *gorm.DB, values interface{}) error {
	//UPDATE article SET ... WHERE id=a.ID AND is_del=0;
	if err:= db.Model(&Article{}).Where("id=? AND is_del=?",a.ID,0).
		Updates(values).Error; err != nil {
			return err
	}
	return nil
}

//获取文章
func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	// SELECT * FROM article WHERE id=a.ID AND state=a.State AND is_del=0
	// ORDER BY ID LIMIT 1;
	db = db.Where("id=? AND state=? AND is_del=?",a.ID,a.State,0)
	//返回按住键排序的第一篇文章
	if err:=db.First(&article).Error; err!=nil && err != gorm.ErrRecordNotFound {
		return article,err
	}
	return article,nil
}

//删除文章
func (a Article) Delete(db *gorm.DB) error {
	//DELETE FROM article WHERE id=a.ID AND is_del=0;
	return db.Where("id=? AND is_del=?",a.ID,0).Delete(&Article{}).Error
}

//文章行信息
//用于在列表中批量显示多个文章的信息
type ArticleRow struct {
	ArticleID uint32
	TagID uint32
	TagName string
	ArticleTitle string
	ArticleDesc string
	CoverImageUrl string
	Content string
}

//统计相同标签ID的文章数
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	// 根据文章标签关联表查询带有标签tagID(未禁用未删除)的文章的数量
	// 使用左外连接保证一定包含article_tag中所有项 ?
	// SELECT count(*) FROM
	// (article_tag AS at LEFT JOIN tag AS t ON at.tag_id=t.id)
	// LEFT JOIN article AS ar ON at.article_id=ar.id
	// WHERE at.tag_id=tagID AND ar.State=a.State AND ar.is_del=0;
	if err := db.Table(ArticleTag{}.TableName()+"AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id=t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"AS ar ON at.article_id=ar.id").
		Where("at.`tag_id`=? AND ar.state=? AND ar.is_del=?",tagID,a.State,0).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count,nil
}

//列出所有相同标签ID的文章(跳过前pageOffset个,取pageSize个)
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{
		"ar.id AS article_id","ar.title AS article_title",
		"ar.desc AS article_desc", "ar.cover_image_url",
		"t.id AS tag_id", "t.name AS tag_name",
	}
	//跳过pageOffset个记录,最多取pageSize个记录
	if pageOffset>=0 && pageSize>0{
		db=db.Offset(pageOffset).Limit(pageSize)
	}
	// SELECT ar.id AS article_id, ar.title AS article_title,
	// ar.desc AS article_desc, ar.cover_image_url, t.id AS tag_id, t.name AS tag_name
	// FROM (article_tag AS at LEFT JOIN tag AS t ON at.tag_id=t.id)
	// LEFT JOIN article AS ar ON at.article_id=ar.id
	// WHERE at.tag_id=tagID AND ar.state=a.State AND ar.is_del=0
	// LIMIT pageSize OFFSET pageOffset;
 	rows,err:=db.Select(fields).Table(ArticleTag{}.TableName()+"AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id=t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id=ar.id").
		Where("at.`tag_id`=? AND ar.state=? AND ar.is_del=?",tagID,a.State,0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
 	//遍历查询结果记录到文章行信息的切片中
	for rows.Next() {
		r:=&ArticleRow{}
		if err := rows.Scan(
			&r.ArticleID,&r.ArticleTitle,&r.ArticleDesc,&r.CoverImageUrl,&r.Content,
			&r.TagID,&r.TagName,
			); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles,nil
}
