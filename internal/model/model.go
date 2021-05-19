package model

import (
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/pkg/setting"
	"time"
)

const (
	StateOpen   = 1
	StateClose = 0
)

//数据表公共部分模型
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedOn  uint32 `json:"created_on"`  //创建时间
	CreatedBy  string `json:"created_by"`  //创建人
	ModifiedOn uint32 `json:"modified_on"` //修改时间
	ModifiedBy string `json:"modified_by"` //修改人
	DeletedOn  uint32 `json:"deleted_on"`   //删除时间
	IsDel      uint8  `json:"is_del"`      //是否删除
}

//新建数据库引擎
func NewDBEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	//注意该部分原书有误
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	)
	db, err := gorm.Open(dbSetting.DBType, s)
	if err != nil {
		return nil, err
	}
	// Debug模式设置数据库记录详细日志
	if global.ServerSetting.RunMode == "Debug" {
		db.LogMode(true)
	}
	//gorm中数据表名默认为结构体名复数形式(结构体User->数据表users)
	//使用该函数转义结构体名时不加上s
	db.SingularTable(true)
	//注册(替换)数据库操作的回调操作
	db.Callback().Create().Replace("gorm:update_time_stamp",
		updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",
		updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",deleteCallback)

	//设置数据库最大空闲连接数和最大连接数
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)
	//添加数据库回调
	otgorm.AddGormCallbacks(db)
	return db, nil
}

//新增操作的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	// 作用域 gorm.Scope 包含操作数据库时的操作信息
	//PS: gorm.Scope在gorm v2即v1.20.0中已经移除
	if scope.HasError() {
		return
	}
	nowTime := time.Now().Unix()
	//判断当前Scope是否包含CreateOn字段
	if createTimeField,ok:=scope.FieldByName("CreatedOn"); ok {
		//CreatedOn字段为空则设置为当前时间
		if createTimeField.IsBlank {
			_ = createTimeField.Set(nowTime)
		}
	}
	//判断是否包含ModifiedOn字段
	if modifyTimeField,ok:=scope.FieldByName("ModifiedOn"); ok {
		//ModifiedOn字段为空则设为当前时间
		if modifyTimeField.IsBlank {
			_ = modifyTimeField.Set(nowTime)
		}
	}
}

//更新操作的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//判断当前操作中是否标识了gorm:update_column属性
	if _,ok := scope.Get("gorm:update_column");!ok{
		//设置ModifiedOn字段
		_ = scope.SetColumn("ModifiedOn",time.Now().Unix())
	}
}

//在非空字符串前添加空格
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " "+str
	}
	return str
}

//删除操作的回调
func deleteCallback(scope *gorm.Scope) {
	if scope.HasError(){
		return
	}
	var extraOption string
	//判断当前操作中是否标识了gorm:delete_option属性
	if str,ok := scope.Get("gorm:delete_option");ok{
		extraOption = fmt.Sprint(str)
	}
	//判断当前是否包含DeletedOn和DsDel字段
	deleteOnField,hasDeletedOnField := scope.FieldByName("DeletedOn")
	isDelField,hasIsDelField:=scope.FieldByName("IsDel")
	if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
		//软删除
		now := time.Now().Unix()
		scope.Raw(fmt.Sprintf(
			"UPDATE %v SET %v=%v,%v=%v%v%v",
			scope.QuotedTableName(),	//当前表名
			scope.Quote(deleteOnField.DBName),
			scope.AddToVars(now),	//添加值
			scope.Quote(isDelField.DBName),
			scope.AddToVars(1),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
			)).Exec()
	}else {
		//硬删除
		scope.Raw(fmt.Sprintf(
			"DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
			)).Exec()
	}
}