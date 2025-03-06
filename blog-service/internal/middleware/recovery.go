package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/blog-service/global"
	"github.com/wushengyouya/blog-service/pkg/app"
	"github.com/wushengyouya/blog-service/pkg/email"
	"github.com/wushengyouya/blog-service/pkg/errcode"
)

// 自定义recovery，全局异常处理
func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)

				// 邮件通知异常
				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}
				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()

			}
		}()
		ctx.Next()
	}
}
