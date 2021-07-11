package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/service"
	"github.com/golang-minibear2333/gin-blog/pkg/app"
	"github.com/golang-minibear2333/gin-blog/pkg/convert"
	"github.com/golang-minibear2333/gin-blog/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

// List tag list
// @tags tag
// @Summary 获取多个标签
// @Produce  json
// @Param token header string true  "用户token" default(debug)
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态，是否启用(0 为禁用、1 为启用)" Enums(0, 1) default(1)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	// 使用 internal/service/tag.go 内的结构体验证ctx请求带上的参数是否合法，并根据地域进行翻译报错
	valid, errs := app.BindAndValidBody(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 获取Service
	svc := service.New(c.Request.Context())
	// 处理分页参数，防止出现异常
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	// 获取标签总数，这里不用传name时可以拉到所有标签，State代表禁用启用
	// 这里查询列表，如果传入标签名就只查询某同名标签（未做唯一性验证）
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	// 获取标签列表
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	// 返回标签列表成规定的格式
	response.ToResponseList(tags, totalRows)
	return
}

// Create
// @tags tag
// @Summary 新增标签
// @Produce  json
// @accept json
// @Param token header string true  "用户token" default(debug)
// @Param data body service.CreateTagRequest true "请求体"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValidBody(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Update tag update
// @tags tag
// @Summary 更新标签
// @Produce  json
// @accept json
// @Param token header string true  "用户token" default(debug)
// @Param id path int true "标签 ID"
// @Param data body service.UpdateTagRequest true "请求体"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValidBody(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Delete tag delete
// @Summary 删除标签
// @tags tag
// @Produce  json
// @Param token header string true  "用户token" default(debug)
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValidBody(c, &param)
	if !valid {
		// TODO 这些代码重复太多次有办法抽出来吗？
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
