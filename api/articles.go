package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/content/models"
	"github.com/provider-go/pkg/ioput"
)

// CreateArticle 增加文章
func CreateArticle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	channelId := ioput.ParamToInt32(json["channelId"])
	title := ioput.ParamToString(json["title"])
	content := ioput.ParamToString(json["content"])
	err := models.CreateArticle(channelId, title, content)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

// DeleteArticle 删除文章
func DeleteArticle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	err := models.DeleteArticle(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

// ListArticle 文章列表
func ListArticle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := ioput.ParamToInt(json["pageSize"])
	pageNum := ioput.ParamToInt(json["pageNum"])
	list, total, err := models.ListArticle(pageSize, pageNum)

	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		ioput.ReturnSuccessResponse(ctx, res)
	}
}

// ViewArticle 查看文章
func ViewArticle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	row, err := models.ViewArticle(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, row)
	}
}
