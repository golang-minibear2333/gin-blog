package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/pkg/version"
)

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// gin 中的Get Set 元数据管理
		c.Set("app_name", version.AppName)
		c.Set("app_version", version.Version)
		c.Next()
	}
}
