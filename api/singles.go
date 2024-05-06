package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/content/models"
	"github.com/provider-go/pkg/output"
)

// CreateSingle 增加文章
func CreateSingle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	title := output.ParamToString(json["title"])
	content := output.ParamToString(json["content"])
	err := models.CreateSingle(title, content)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

// DeleteSingle 删除文章
func DeleteSingle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	err := models.DeleteSingle(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

// ListSingle 文章列表
func ListSingle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := output.ParamToInt(json["pageSize"])
	pageNum := output.ParamToInt(json["pageNum"])
	list, total, err := models.ListSingle(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}
}

// ViewSingle 查看文章
func ViewSingle(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	row, err := models.ViewSingle(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}
}
