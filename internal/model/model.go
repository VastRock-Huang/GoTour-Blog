package model

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


