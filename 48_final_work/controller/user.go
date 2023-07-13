package controller

import (
	"errors"
	"fmt"
	"go_web_demo/48_final_work/dao/mysqlx"
	"go_web_demo/48_final_work/logic"
	"go_web_demo/48_final_work/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var u = new(models.SignUpStruct)
	// 1. 接受数据，检查数据
	if err := c.ShouldBindJSON(u); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": errs.Translate(trans),
		//})
		return
	}
	//// 业务代码，比如 密码和确认密码必须相同，三个值不能为空
	//if len(u.Username) == 0 || len(u.Password) == 0 || len(u.Repassword) == 0 {
	//	zap.L().Error("参数不能为空")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "参数不能为空",
	//	})
	//	return
	//}
	//if u.Password != u.Repassword {
	//	zap.L().Error("密码和确认密码必须相同")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "密码和确认密码必须相同",
	//	})
	//	return
	//}
	// 2. 注册
	fmt.Println(u.Username, u.Password, u.Repassword)
	err := logic.Signup(u)
	if err != nil {
		zap.L().Error("signup failed", zap.Error(err))
		if errors.Is(err, mysqlx.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Error(),
		//})
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回相应
	ResponseSuccess(c, nil)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "success",
	//})
}

func LoginHandler(c *gin.Context) {
	u := new(models.LoginStruct)
	// 1. 接受数据，检查数据
	if err := c.ShouldBindJSON(u); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": errs.Translate(trans),
		//})
		return
	}

	// 登录账户业务
	token, err := logic.Login(u)
	if err != nil {
		if errors.Is(err, mysqlx.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysqlx.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Error(),
		//})
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "success",
	//})
}
