package logic

import (
	"go_web_demo/48_final_work/dao/mysqlx"
	"go_web_demo/48_final_work/models"
	"go_web_demo/48_final_work/pkg/jwt"
	"go_web_demo/48_final_work/pkg/snowflake"
)

func Signup(u *models.SignUpStruct) (err error) {
	// 1. 检查用户是否已经存在
	if err = mysqlx.CheckUser(u); err != nil {
		return
	}
	// 2. 插入用户
	id := snowflake.GetGenID()
	user := &models.UserStruct{
		UserID:   id,
		Username: u.Username,
		Password: u.Password,
	}
	err = mysqlx.InsertUser(user)
	return err
}

func Login(u *models.LoginStruct) (token string, err error) {
	user := &models.UserStruct{
		Username: u.Username,
		Password: u.Password,
	}
	if err = mysqlx.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.Username, user.UserID)
}
