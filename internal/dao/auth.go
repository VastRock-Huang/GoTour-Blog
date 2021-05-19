package dao

import "github.com/vastrock-huang/gotour-blogservice/internal/model"

//给定appKey和appSecret从数据库中获取认证信息
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
