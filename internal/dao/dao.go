package dao

import "github.com/jinzhu/gorm"

//数据库访问对象DAO
type Dao struct {
	engine *gorm.DB
}

//新建DAO
func New(engine *gorm.DB) *Dao {
	return &Dao{
		engine: engine,
	}
}


