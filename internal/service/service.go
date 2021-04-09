package service
//服务层请求处理
//位于DAO层之上,进行参数验证后将参数传递给DAO

import (
	"context"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/internal/dao"
)

//请求服务结构体
type Service struct {
	ctx context.Context	//上下文
	dao *dao.Dao	//数据库访问对象
}

//新建请求服务
func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
	}
	svc.dao=dao.New(global.DBEngine)
	return svc
}