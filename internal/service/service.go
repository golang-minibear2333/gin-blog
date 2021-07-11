package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

// New 所有Service都走这里的逻辑
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	// 创建dao层逻辑，并传入SQL链路追踪（数据库连接的上下文注册），官方建议上下文参数作为第一个参数
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
