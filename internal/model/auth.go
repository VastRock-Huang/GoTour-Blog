package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type AuthSwagger struct {
	Token string `json:"token"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

//获取认证记录
func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	// SELECT * FROM auth WHERE app_key=? AND app_secret=? AND is_del=0 LIMIT 1;
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?",
		a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error

	if err != nil {
		return auth,err
	}
	////若有错但有记录则返回
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return auth, err
	//}
	return auth, nil
}
