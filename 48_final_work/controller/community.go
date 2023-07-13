package controller

import (
	"go_web_demo/48_final_work/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	// 1. 查询到所有社区的id 和名字
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	// 解析并验证参数
	uid := c.Param("id")
	id, err := strconv.ParseInt(uid, 10, 46)
	if err != nil {
		zap.L().Error("parse failed,", zap.Error(err))
		return
	}
	// 2. 获取用户数据
	communityDetail, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, communityDetail)
}
