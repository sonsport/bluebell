package controller

import (
	"go_web_demo/48_final_work/logic"
	"go_web_demo/48_final_work/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func PostVote(c *gin.Context) {
	// 接受和检查数据
	vote := new(models.VoteData)
	userID, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get current user failed", zap.Error(err))
		ResponseError(c, CodeNeedAuth)
		return
	}
	if err = c.ShouldBindJSON(vote); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := errs.Translate(trans)
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//传入redis
	err = logic.PostVote(userID, vote)
	if err != nil {
		zap.L().Error("logic post vote failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//  返回响应
	ResponseSuccess(c, nil)

}
