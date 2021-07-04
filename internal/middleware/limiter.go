package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/pkg/app"
	"github.com/golang-minibear2333/gin-blog/pkg/errcode"
	"github.com/golang-minibear2333/gin-blog/pkg/limiter"
)

// RateLimiter 限流器
// 入参应该为 LimiterIface 接口类型，这样子的话只要符合该接口类型的具体限流器实现都可以传入并使用
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// 占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，如果没有可用的令牌，将会返回 0，也就是已经超出配额了
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				// 如果没有可用的令牌，返回 errcode.TooManyRequest 状态告诉客户端需要减缓并控制请求速度。
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
