package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 令牌限流
type MethodLimiter struct {
	*Limiter
}

func NewMetHodLimiter() LimiterIface {
	return MethodLimiter{Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}}
}

// 根据uri生成键 api/v1/tags?id=1 => api/v1/tags 使不同参数的相同路径共享同一个限流配置
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]

}
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}
func (l MethodLimiter) AddBuckets(rules ...LimiterBuckerRule) LimiterIface {
	for _, rule := range rules {

		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(
				rule.FillInterval, // 令牌补充间隔时间
				rule.Capacity,     // 桶的最大容量
				rule.Quantum,      // 每次补充的令牌数
			)
		}
	}
	return l
}
