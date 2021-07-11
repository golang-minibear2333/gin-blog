package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/service"
	"github.com/golang-minibear2333/gin-blog/pkg/app"
	"github.com/golang-minibear2333/gin-blog/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	// 此结构体内限定了app_key 和 app_secrect是必填字段
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	// 校验及获取入参,绑定获取到的 app_key 和 app_secrect
	valid, errs := app.BindAndValidHeader(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	// 进行数据库查询，检查认证信息是否存在
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	// 若存在则进行 Token 的生成并返回。
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
