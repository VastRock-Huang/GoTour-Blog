package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/service"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
	"github.com/vastrock-huang/gotour-blogservice/pkg/convert"
	"github.com/vastrock-huang/gotour-blogservice/pkg/errcode"
	"github.com/vastrock-huang/gotour-blogservice/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

//上传文件
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)	//响应体
	//读取file字段的上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		errRsp :=errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	//文件头信息为空或文件类型有误则返回
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	//上传文件
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c,"svc.UploadFile err: %v", err)
		errRsp:=errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	//返回上传后文件对应的Url
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
