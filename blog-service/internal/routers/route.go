package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wushengyouya/blog-service/internal/routers/api/v1"
)

func NewRouters() *gin.Engine {

	article := v1.NewArticle()
	tag := v1.NewTag()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 创建路由组
	apiv1 := r.Group("/api/v1")
	{
		// 标记路由
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// 文章路由
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}
	return r
}
