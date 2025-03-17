package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/blog-service/global"
	"github.com/wushengyouya/blog-service/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, nil
	}

	return w.ResponseWriter.Write(p)
}

// AccessLog 访问日志中间件
func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyWriter

		beginTime := time.Now().Unix()

		ctx.Next()

		endTime := time.Now().Unix()

		var requestData string

		// TODO:存在问题
		if ctx.Request.Method == http.MethodGet {
			requestData = ctx.Request.URL.Query().Encode()
		} else {
			contentType := ctx.ContentType()
			switch contentType {
			case "application/x-www-form-urlencoded", "multipart/form-data":
				_ = ctx.Request.ParseForm()
				requestData = ctx.Request.PostForm.Encode()
			case "application/json":
				bodyBytes, _ := ctx.GetRawData()
				requestData = string(bodyBytes)

			}
		}
		// 此处读取不到值
		fields := logger.Fields{
			"request":  requestData,
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			ctx.Request.Method, bodyWriter.Status(), beginTime, endTime)

	}
}
