package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-minibear2333/gin-blog/docs"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/middleware"
	"github.com/golang-minibear2333/gin-blog/internal/routers/api"
	v1 "github.com/golang-minibear2333/gin-blog/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// 类似过滤器
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	// 访问 /swagger/index.html 可以查看效果
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") // 重定向
	})
	article := v1.NewArticle()
	tag := v1.NewTag()
	// 上传文件功能
	upload := api.NewUpload()
	// curl -X POST http://127.0.0.1:8000/upload/file -F file=@/Users/xxx/Downloads/golang.png  -F type=1
	r.POST("/upload/file", upload.UploadFile)
	// 提供静态文件目录访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	// JWT 权限校验
	r.POST("/auth", api.GetAuth)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		tags := v1.Group("/tags")
		{
			tags.POST("", tag.Create)
			v1.DELETE(":id", tag.Delete)
			v1.PUT(":id", tag.Update)
			v1.PATCH(":id/state", tag.Update)
			v1.GET("", tag.List)
		}
		articles := v1.Group("/articles")
		{
			articles.POST("/articles", article.Create)
			articles.DELETE("/articles/:id", article.Delete)
			articles.PUT("/articles/:id", article.Update)
			articles.PATCH("/articles/:id/state", article.Update)
			articles.GET("/articles/:id", article.Get)
			articles.GET("/articles", article.List)
		}

	}
	return r
}
