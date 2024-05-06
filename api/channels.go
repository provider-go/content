package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/content/models"
	"github.com/provider-go/pkg/output"
)

// CreateChannel 增加频道
func CreateChannel(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pid := output.ParamToInt32(json["pid"])
	channelName := output.ParamToString(json["channelName"])
	err := models.CreateChannel(pid, channelName)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

// DeleteChannel 删除频道
func DeleteChannel(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	err := models.DeleteChannel(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

// ListChannel 频道列表
func ListChannel(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := output.ParamToInt(json["pageSize"])
	pageNum := output.ParamToInt(json["pageNum"])
	list, total, err := models.ListChannel(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}
}
