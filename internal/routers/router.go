package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-minibear2333/gin-blog/docs"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/middleware"
	"github.com/golang-minibear2333/gin-blog/internal/routers/api"
	v1 "github.com/golang-minibear2333/gin-blog/internal/routers/api/v1"
	"github.com/golang-minibear2333/gin-blog/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 限流模块
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		// 访问记录日志
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	// 类似过滤器
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
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
			tags.DELETE(":id", tag.Delete)
			tags.PUT(":id", tag.Update)
			tags.PATCH(":id/state", tag.Update)
			tags.GET("", tag.List)
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
