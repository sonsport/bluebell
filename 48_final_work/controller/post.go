package controller

import (
	"fmt"
	"go_web_demo/48_final_work/logic"
	"go_web_demo/48_final_work/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	fmt.Println("****************end")
	var err error
	// 接收和检查数据
	post := new(models.Post)
	if err = c.ShouldBindJSON(post); err != nil {
		zap.L().Error("ShouldBindJSON failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Println("****************start")
	fmt.Println(post.CommunityID)
	fmt.Println(post.Content)
	fmt.Println(post.Title)
	// 2. 写入数据库
	post.AuthorID, _, err = GetCurrentUser(c)
	if err != nil {
		zap.L().Error("GetCurrentUser failed", zap.Error(err))
		ResponseError(c, CodeNeedAuth)
		return
	}
	if err = logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost(post) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func GetDetailPost(c *gin.Context) {
	// 1. 接收并检查数据
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		zap.L().Error("parse failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Println(id)
	// 2. 获取到数据库信息
	post, err := logic.GetDetailPost(id)
	// 最好用名字 GetPostByID
	if err != nil {
		zap.L().Error("get post failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, post)
}

func GetPost(c *gin.Context) {
	site, page, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error("get page info failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询所有数据
	post, err := logic.GetPost(site, page)
	if err != nil {
		zap.L().Error("logic get post failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, post)
}

func GetPostOrder(c *gin.Context) {
	queryStr := &models.GetPostOder{
		Order: "score",
		Site:  10,
		Page:  1,
	}
	if err := c.ShouldBindQuery(queryStr); err != nil {
		zap.L().Error("should bind query failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostByOrder(queryStr)
	if err != nil {
		zap.L().Error("get post by order failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostOrder2(c *gin.Context) {
	queryStr := &models.GetPostOder{
		Order: "score",
		Site:  10,
		Page:  1,
	}
	if err := c.ShouldBindQuery(queryStr); err != nil {
		zap.L().Error("should bind query failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostByOrder2(queryStr)
	if err != nil {
		zap.L().Error("get post by order failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetCommunityPostOrder(c *gin.Context) {
	queryStr := &models.CommunityPostOrder{
		Order: "score",
		Site:  10,
		Page:  1,
	}
	if err := c.ShouldBindQuery(queryStr); err != nil {
		zap.L().Error("should bind query failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityPostOrder(queryStr)
	if err != nil {
		zap.L().Error("logic get community post order failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetDetail(c *gin.Context) {
	_, username, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get current user failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetTableDetail(username)
	if err != nil {
		zap.L().Error("logic get table detail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
