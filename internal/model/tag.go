package model
//数据库模型层访问

import (
	"github.com/jinzhu/gorm"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
)

//标签信息
type Tag struct {
	*Model
	Name string `json:"name"`	//标签名
	State uint8 `json:"state"`	//标签状态
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

//获取标签表名
func (t Tag) TableName() string {
	return "blog_tag"
}

//统计标签数
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=?",t.Name)	//选择名为给定标签的记录
	}
	db = db.Where("state=?",t.State)	//筛选状态
	//SELECT count(*) FROM tag
	// WHERE id=t.ID AND is_del=0 AND name=t.Name AND state=t.State;
	err := db.Model(&t).Where("is_del=?",0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count,nil
}

//列出所有标签
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		//跳过pageOffset个记录,限制最多获取pageSize个记录
		//SELECT * FROM tag OFFSET=pageOffset LIMIT pageSize;
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db=db.Where("name=?",t.Name)
	}
	db=db.Where("state=?",t.State)
	//SELECT * FROM tag WHERE is_del=0 AND name=t.State AND name=t.Name;
	//结果存入tags中
	err = db.Where("is_del=?",0).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags,nil
}

//创建一个标签
func (t Tag) Create(db *gorm.DB) error {
	//INSERT INTO tag VALUES(...)
	//将t作为一条记录插入数据库
	return db.Create(&t).Error
}

//更新标签
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//UPDATE tag SET ... WHERE id=t.ID AND is_del=0;
	//此处使用Updates更新映射values中的字段,以防止将state置为0后而未被gorm修改
	if err := db.Model(&Tag{}).Where("id=? AND is_del=?",t.ID,0).
		Updates(values).Error; err != nil {
			return err
	}
	return nil
	//db = db.Model(&Tag{}).Where("id=? AND is_del=?",t.ID,0)
	//return db.Update(t).Error
}

//删除标签
func (t Tag) Delete(db *gorm.DB) error {
	//DELETE FROM tag WHERE id=t.ID AND is_del=0
	return db.Where("id=? AND is_del=?",t.ID,0).Delete(&Tag{}).Error
	//PS:对于上述gorm语句和以下操作代码的
	//因为Delete()函数参数若为具体的对象,则会增加其主键作为条件进行删除
	//db.Where("id=? AND is_del=?",t.ID,0).Delete(&t).Error		//原代码
	//db.Where("is_del=?",0).Delete(&t).Error
}