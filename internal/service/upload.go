package service

import (
	"errors"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

//上传文件
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File,
	fileHeader *multipart.FileHeader) (*FileInfo,error) {
	fileName := upload.GetFileName(fileHeader.Filename)	//加密后的文件名
	//检测拓展名是否支持
	if !upload.CheckContainExt(fileType,fileName) {
		return nil,errors.New("file suffix is not supported")
	}
	uploadSavePath := upload.GetSavePath()
	//检测文件是否存在,不存在则创建
	if upload.CheckSavePathNotExist(uploadSavePath) {
		if err:=upload.CreateSavePath(uploadSavePath,os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	//检测文件是否超限
	if upload.CheckMaxSize(fileType,file) {
		return nil,errors.New("exceeded maximum file limit")
	}
	//检测文件是否权限不足
	if upload.CheckNotPermission(uploadSavePath) {
		return nil,errors.New("insufficient file permissions")
	}
	dst := uploadSavePath + "/" +fileName
	//保存文件
	if err := upload.SaveFile(fileHeader,dst); err != nil {
		return nil, err
	}
	accessUrl:=global.AppSetting.UploadServerUrl+"/"+fileName
	return &FileInfo{
		Name: fileName,
		AccessUrl: accessUrl,
	},nil
}