package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// ContextTimeOut 访问超时中间件
func ContextTimeOut(t time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx.Request.Context(), t)
		defer cancel()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
