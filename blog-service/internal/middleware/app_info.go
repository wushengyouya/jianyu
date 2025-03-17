package middleware

import "github.com/gin-gonic/gin"

// AppInfo 应用信息中间件
func AppInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("app_name", "blog-service")
		ctx.Set("app_version", "1.0.0")
		ctx.Next()
	}
}
