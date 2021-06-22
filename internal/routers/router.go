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

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
