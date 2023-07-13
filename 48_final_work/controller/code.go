package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedAuth
	CodeInvalidToken
)

var codeMsgmap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",

	CodeNeedAuth:     "需要认证",
	CodeInvalidToken: "认证错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgmap[c]
	if !ok {
		msg = codeMsgmap[CodeServerBusy]
	}
	return msg
}
