package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserID = "userID"
const CtxUsername = "username"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (userID int64, Username string, err error) {
	uid, ok := c.Get(CtxUserID)
	uname, uok := c.Get(CtxUsername)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	if !uok {
		err = ErrorUserNotLogin
		return
	}
	Username, uok = uname.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64, error) {
	//接收和检查数据
	sSite := c.Query("site")
	sPage := c.Query("page")
	var (
		site int64
		page int64
		err  error
	)
	site, err = strconv.ParseInt(sSite, 10, 64)
	if err != nil {
		site = 10
	}
	page, err = strconv.ParseInt(sPage, 10, 64)
	if err != nil {
		page = 1
	}
	return site, page, nil
}
