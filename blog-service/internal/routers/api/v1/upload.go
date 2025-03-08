package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/blog-service/global"
	"github.com/wushengyouya/blog-service/internal/services"
	"github.com/wushengyouya/blog-service/pkg/app"
	"github.com/wushengyouya/blog-service/pkg/convert"
	"github.com/wushengyouya/blog-service/pkg/errcode"
	"github.com/wushengyouya/blog-service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// 单文件上传中间件
// TODO:多文件上传
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")

	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := services.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})

}
