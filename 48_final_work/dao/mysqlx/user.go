package mysqlx

import (
	"crypto/md5"
	sql2 "database/sql"
	"encoding/hex"
	"errors"
	"go_web_demo/48_final_work/models"
)

const secret = "wanoqingyun"

var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码不正确")
)

func CheckUser(u *models.SignUpStruct) (err error) {
	sql := `select count(user_id) from user where username = ?`
	count := new(int)
	if err = db.Get(count, sql, u.Username); err != nil {
		return
	}
	if *count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(u *models.UserStruct) (err error) {
	u.Password = encryPassword(u.Password)
	sql := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sql, u.UserID, u.Username, u.Password)
	return
}

func encryPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(u *models.UserStruct) (err error) {
	password := encryPassword(u.Password)
	sql := `select user_id,username,password from user where username = ?`
	err = db.Get(u, sql, u.Username)
	if err == sql2.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if password != u.Password {
		return ErrorInvalidPassword
	}
	return
}
