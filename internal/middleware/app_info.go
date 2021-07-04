package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// gin 中的Get Set 元数据管理
		c.Set("app_name", "gin-blog")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
