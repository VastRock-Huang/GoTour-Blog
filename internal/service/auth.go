package service

import "errors"

//认证信息请求
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

//认证信息校验
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return err
	}
	//若能从数据库中得到对应的认证信息则成功
	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}