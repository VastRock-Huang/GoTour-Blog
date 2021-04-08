package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
)

type Article struct {
	
}


func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)  {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (a Article) List(c *gin.Context) {
	
}

func (a Article) Create(c *gin.Context)  {
	
}

func (a Article) Update(c *gin.Context)  {
	
}

func (a Article) Delete(c *gin.Context)  {
	
}
