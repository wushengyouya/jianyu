package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wushengyouya/blog-service/docs"
	"github.com/wushengyouya/blog-service/internal/middleware"
	v1 "github.com/wushengyouya/blog-service/internal/routers/api/v1"
)

func NewRouters() *gin.Engine {

	article := v1.NewArticle()
	tag := v1.NewTag()

	// 初始化engine
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// TODO:放在中间件每次请求注册会Panic,后续调整到init函数中只初始化一次
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 创建路由组
	apiv1 := r.Group("/api/v1")
	{
		// 标记路由
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags", tag.Delete)
		apiv1.PUT("/tags", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// 文章路由
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles", article.Delete)
		apiv1.PUT("/articles", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}
	return r
}
