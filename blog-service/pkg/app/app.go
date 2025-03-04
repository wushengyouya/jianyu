package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/blog-service/pkg/errcode"
)

// 响应库
type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRow int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data any) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list any, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:     GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRow: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}

func (r *Response) ToSuccessResponse(msg any) {
	response := gin.H{"success": "ok", "msg": msg}
	r.Ctx.JSON(http.StatusOK, response)
}
