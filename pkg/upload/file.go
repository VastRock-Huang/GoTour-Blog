package upload

import (
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int	//文件类型

//图片文件类型
const TypeImage FileType = iota + 1

//获取MD5加密后的文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)		//拓展名
	fileName := strings.TrimSuffix(name, ext)	//去除拓展名后的文件名
	fileName = util.EncodeMD5(fileName)		//MD5加密
	return fileName + ext
}

//获取文件拓展名
func GetFileExt(name string) string {
	return path.Ext(name)
}

//获取文件保存路径
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检测文件保存路径是否不存在
func CheckSavePathNotExist(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//检测是否支持当前文件的后缀名
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)	//后缀名
	switch t {
	case TypeImage:
		//遍历设置中允许的后缀名,判断有无相同的
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToLower(allowExt) == strings.ToLower(ext) {
				return true
			}
		}
	}
	return false
}

//检测文件是否超出文件大小限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)	//读取文件内容
	size := len(content)	//获取文件大小
	switch t {
	case TypeImage:
		//判断大小是否超限
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

//检测文件权限是否不足够
func CheckNotPermission(dst string) bool {
	_, err:=os.Stat(dst)
	return os.IsPermission(err)
}

//创建保存上传文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	if err := os.MkdirAll(dst,perm); err!=nil{
		return err
	}
	return nil
}

//保存上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src,err:=file.Open()	//打开源文件
	if err != nil {
		return err
	}
	defer src.Close()

	out,err := os.Create(dst) //创建新文件
	if err != nil {
		return err
	}
	defer out.Close()
	_,err=io.Copy(out,src)	//文件拷贝
	return err
}
