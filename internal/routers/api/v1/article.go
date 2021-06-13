package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/pkg/app"
	"github.com/golang-minibear2333/gin-blog/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)    {
}
func (a Article) List(c *gin.Context)   {
	// 接口测试 http://localhost:8000/api/v1/article
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
