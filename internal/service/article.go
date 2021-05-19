package service

import (
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
)

//请求参数结构体

//获取文章请求参数
type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//列出文章请求参数
type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//创建文章请求参数
type CreateArticleRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	Title string `form:"title" binding:"required,min=2,max=100"`
	Desc string `form:"desc" binding:"required,min=2,max=255"`
	Content string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

//更新文章请求参数
type UpdateArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	Title string `form:"title" binding:"max=100"`
	Desc string `form:"desc" binding:"max=255"`
	Content string `form:"content" binding:"max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State uint8 `form:"state" binding:"oneof=0 1"`
}

//删除文章请求参数
type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//带标签信息的文章
type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

//获取文章
func (svc *Service) GetArticle(param *GetArticleRequest) (*Article, error) {
	//获取给定文章id
	article,err:=svc.dao.GetArticle(param.ID,param.State)
	if err != nil {
		return nil, err
	}
	//由文章id获取一条文章标签记录
	articleTag,err := svc.dao.GetArticleTagByArticleID(article.ID)
	if err != nil {
		return nil, err
	}
	//给定标签id获取一条标签
	tag,err:= svc.dao.GetTag(articleTag.TagID,model.StateOpen)
	if err != nil {
		return nil, err
	}
	//返回带标签信息的文章
	return &Article{
		ID: article.ID,
		Title: article.Title,
		Desc: article.Desc,
		Content: article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State: article.State,
		Tag: &tag,
	},nil
}

//获取给定标签id的文章列表
func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	//给定标签id统计该标签的文章数
	articleCount,err:=svc.dao.CountArticleListByTagID(param.TagID,param.State)
	if err != nil {
		return nil, 0, err
	}
	//根据标签id获取文章行信息(给定页号和页大小)
	articles,err := svc.dao.GetArticleListByTagID(param.TagID,param.State,pager.Page,pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	//获取带标签信息的文章列表
	var articleList []*Article
	for _,article := range articles {
		articleList = append(articleList, &Article{
			ID:article.ArticleID,
			Title: article.ArticleTitle,
			Desc: article.ArticleDesc,
			Content: article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag: &model.Tag{
				Name: article.TagName,
				Model: &model.Model{
					ID:article.TagID,
				},
			},
		})
	}
	return articleList,articleCount,nil
}

//创建文章
func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	//创建文章
	article,err:=svc.dao.CreateArticle(param.Title,param.Desc,param.Content,
		param.CoverImageUrl,param.State,param.CreatedBy)
	if err != nil {
		return err
	}
	//创建文章对应的文章标签信息
	err = svc.dao.CreateArticleTag(article.ID,param.TagID,param.CreatedBy)
	if err != nil {
		return err
	}
	return err
}

//更新文章
func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	if err:=svc.dao.UpdateArticle(param.ID,param.Title,param.Desc,param.Content,
		param.CoverImageUrl,param.State,param.ModifiedBy);err!=nil {
		return err
	}
	//更新文章对应的文章标签信息
	if err:=svc.dao.UpdateArticleTag(param.ID,param.TagID,
		param.ModifiedBy);err!=nil{
		return err
	}
	return nil
}

//删除文章
func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	if err:= svc.dao.DeleteArticle(param.ID); err!=nil{
		return err
	}
	//删除文章对应的标签信息
	if err:=svc.dao.DeleteArticleTag(param.ID); err!=nil {
		return err
	}
	return nil
}