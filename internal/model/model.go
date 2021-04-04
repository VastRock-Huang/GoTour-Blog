package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/pkg/setting"
)

//数据表公共部分模型
type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreateOn uint32 `json:"create_on"`		//创建时间
	CreateBy string `json:"create_by"`		//创建人
	ModifiedOn uint32 `json:"modified_on"`	//修改时间
	ModifiedBy	string `json:"modified_by"`	//修改人
	DeleteOn uint32 `json:"delete_on"`		//删除时间
	IsDel uint8 `json:"is_del"`				//是否删除
}

func NewDBEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err:= gorm.Open(dbSetting.DBType,
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
		)
	if err != nil{
		return nil, err
	}
	// Debug模式设置数据库记录详细日志
	if global.ServerSetting.RunMode=="Debug"{
		db.LogMode(true)
	}
	//gorm中数据表名默认为结构体名复数形式(结构体User->数据表users)
	//使用该函数转义结构体名时不加上s
	db.SingularTable(true)
	//设置数据库最大空闲连接数和最大连接数
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db,nil
}


