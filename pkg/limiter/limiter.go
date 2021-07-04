package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// LimiterIface 用于定义当前限流器所必须要的方法
// go 语言中的面向对象
// 不同的接口可能要不同的限流器，但要统一限流器内要实现的方法
type LimiterIface interface {
	// Key 获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	// GetBucket 获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// AddBuckets 新增多个令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

// Limiter 存储令牌桶与键值对名称的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// LimiterBucketRule 存储令牌桶的一些相应规则属性
type LimiterBucketRule struct {
	// 自定义键值对名称
	Key string
	// 间隔多久时间放 N 个令牌
	FillInterval time.Duration
	// 令牌桶的容量
	Capacity int64
	// 每次到达间隔时间后所放的具体令牌数量
	Quantum int64
}
