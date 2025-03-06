package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/blog-service/pkg/app"
	"github.com/wushengyouya/blog-service/pkg/errcode"
	"github.com/wushengyouya/blog-service/pkg/limiter"
)

// 限流中间件
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}
