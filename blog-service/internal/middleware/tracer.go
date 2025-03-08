package middleware

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/blog-service/pkg/tracer"
	"go.opentelemetry.io/otel"
)

// TODO:链路追踪中间件
var (
	otelTracer = otel.Tracer("roll")
)

func Tracing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置OpenTelemetry。
		serviceName := "dice"
		serviceVersion := "0.1.0"
		otelShutdown, err := tracer.SetupOTelSDK(ctx, serviceName, serviceVersion)
		if err != nil {
			//return err
		}

		_, span := otelTracer.Start(ctx.Request.Context(), "roll") // 开始 span
		defer span.End()                                           // 结束 span

		// 适当处理关闭，以避免泄漏。
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
		ctx.Next()
	}
}
