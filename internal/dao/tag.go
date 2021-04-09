package dao

import (
	"github.com/vastrock-huang/gotour-blogservice/internal/model"
	"github.com/vastrock-huang/gotour-blogservice/pkg/app"
)

//给定标签名和其状态统计标签数
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag:=model.Tag{
		Name: name,
		State: state,
	}
	return tag.Count(d.engine)
}

//给定标签名,状态,页号和页大小获得标签列表
func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag,error) {
	tag:=model.Tag{
		Name: name,
		State: state,
	}
	pageOffset:=app.GetPageOffset(page,pageSize)
	return tag.List(d.engine,pageOffset,pageSize)
}

//给定标签名,状态和创建人创建标签
func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag:=model.Tag{
		Name: name,
		State: state,
		Model:&model.Model{
			CreatedBy: createdBy,
		},
	}
	return tag.Create(d.engine)
}

//给定标签id,标签名,状态,修改人修改标签
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag:=model.Tag{
		Model: &model.Model{
			ID:id,
		},
	}
	//指定具体修改的字段
	//此处选择额外传递是由于state置0后,gorm会将其作为默认不修改的字段而不会修改数据库
	//使用映射values来传递来确定具体要修改的字段
	values := map[string]interface{}{
		"state":state,
		"modified_by":modifiedBy,
	}
	if name != "" {
		values["name"]=name
	}
	return tag.Update(d.engine,values)
}

//给定标签id删除标签
func (d *Dao) DeleteTag(id uint32) error {
	tag:=model.Tag{
		Model:&model.Model{
			ID: id,
		},
	}
	return tag.Delete(d.engine)
}


