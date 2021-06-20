package service

import (
	"github.com/golang-minibear2333/gin-blog/internal/model"
	"github.com/golang-minibear2333/gin-blog/pkg/app"
)

// CountTagRequest 定义了 Request 结构体作为接口入参的基准，而本项目由于并不会太复杂，所以直接放在了 service 层中便于使用
// 若后续业务不断增长，程序越来越复杂，service 也冗杂了，可以考虑将抽离一层接口校验层，便于解耦逻辑。
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" json:"name" binding:"required,min=2,max=100" `            // 标签名称
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"` //创建者
	State     uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`              //状态，是否启用(0 为禁用、1 为启用)
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`                  // 标签id
	Name       string `form:"name" binding:"max=100"`                       // 标签名称
	State      uint8  `form:"state" binding:"oneof=0 1"`                    //状态，是否启用(0 为禁用、1 为启用)
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"` // 修改者
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// 如下 在应用分层中，service 层主要是针对业务逻辑的封装，如果有一些业务聚合和处理可以在该层进行编码，同时也能较好的隔离上下两层的逻辑

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
