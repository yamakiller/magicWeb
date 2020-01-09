package code

import (
	"fmt"

	"github.com/yamakiller/magicWeb/library/message"
)

//SpawnResponse doc
//ummary create success  response
func SpawnResponse(data interface{}) message.Response {
	return message.Response{Code: 0, Data: data}
}

//SpawnErrParam doc
//Summary create param error response json
//Return (message.Response)
func SpawnErrParam() message.Response {
	return message.Response{Code: ErrParam, Data: ErrMsgParam}
}

//SpawnErrSystem doc
//Summary create system error response json
//Return (message.Response)
func SpawnErrSystem() message.Response {
	return message.Response{Code: ErrSystem, Data: ErrMsgSystem}
}

//SpawnErrSystemMsg doc
//Summary create system error message response json
//Return (message.Response)
func SpawnErrSystemMsg(msg interface{}) message.Response {
	return message.Response{Code: ErrSystem, Data: msg}
}

//SpawnErrDbAbnormal doc
//Summary create db/mysql abnormal error response json
//Return (message.Response)
func SpawnErrDbAbnormal() message.Response {
	return message.Response{Code: ErrDbAbnormal, Data: ErrMsgDbAbnormal}
}

//SpawnErrCaptchaFormat doc
//Summary create captcha format error response json
//Return (message.Response)
func SpawnErrCaptchaFormat(msg string) message.Response {
	return message.Response{Code: ErrCaptchaFormat, Data: fmt.Sprintf("%s:%s", ErrMsgCaptchaFormat, msg)}
}

//SpawnErrUserNameFormat doc
//Summary create username format error response json
//Return (message.Response)
func SpawnErrUserNameFormat() message.Response {
	return message.Response{Code: ErrUserNameFormat, Data: ErrMsgUserNameFormat}
}

//SpawnErrPwdFormat doc
//Summary create password format error response json
//Return (message.Response)
func SpawnErrPwdFormat() message.Response {
	return message.Response{Code: ErrPwdFormat, Data: ErrMsgPwdFormat}
}

//SpawnErrUserName doc
//Summary create user name error response json
//Return (message.Response)
func SpawnErrUserName() message.Response {
	return message.Response{Code: ErrUserName, Data: ErrMsgUserName}
}

//SpawnErrPwd doc
//Summary create password error response json
//Return (message.Response)
func SpawnErrPwd() message.Response {
	return message.Response{Code: ErrPwd, Data: ErrMsgPwd}
}

//SpawnErrPwdAgin doc
//Method SpawnErrPwdAgin @Summary create password again error response json
//Return (message.Response)
func SpawnErrPwdAgin() message.Response {
	return message.Response{Code: ErrPwdAgain, Data: ErrMsgPwdAgain}
}

//SpawnErrTokenNot doc
//Method SpawnErrTokenNot @Summary create authorization token error response json
//Return (message.Response)
func SpawnErrTokenNot() message.Response {
	return message.Response{Code: ErrTokenNot, Data: ErrMsgTokenNot}
}

//SpawnErrTokenInvalid doc
//Method SpawnErrTokenInvalid @Summary create authorization token invalid error response json
//Return (message.Response)
func SpawnErrTokenInvalid() message.Response {
	return message.Response{Code: ErrTokenInvalid, Data: ErrMsgTokenInvalid}
}

//SpawnErrOnlineUserNot doc
//Method SpawnErrOnlineUserNot @Summary create user unonline error response json
//Return (message.Response)
func SpawnErrOnlineUserNot() message.Response {
	return message.Response{Code: ErrOnlineUserNot, Data: ErrMsgOnlineUserNot}
}

//SpawnErrNeedPerm doc
//Method SpawnErrNeedPerm @Summary create need sufficient permissions error  response json
//Return (message.Response)
func SpawnErrNeedPerm() message.Response {
	return message.Response{Code: ErrNeedPerm, Data: ErrMsgNeedPerm}
}

//SpawnErrNeedGenerateCaptcha doc
//Method SpawnErrNeedGenerateCaptcha @Summary create captcha not grenerate error  response json
//Return (message.Response)
func SpawnErrNeedGenerateCaptcha() message.Response {
	return message.Response{Code: ErrNeedGenerateCaptcha, Data: ErrMsgNeedGenerateCaptcha}
}

//SpawnErrCaptcha doc
//Summary create captcha error  response json
//Method SpawnErrCaptcha
//Return (message.Response)
func SpawnErrCaptcha() message.Response {
	return message.Response{Code: ErrCaptcha, Data: ErrMsgCaptcha}
}

//SpawnErrConfigNot doc
//Summary create config data non-existent error response json
//Method SpawnErrConfigNot
//Return (message.Response)
func SpawnErrConfigNot() message.Response {
	return message.Response{Code: ErrConfigNot, Data: ErrMsgConfigNot}
}

//SpawnErrUserExitis doc
func SpawnErrUserExitis() message.Response {
	return message.Response{Code: ErrUserExitis, Data: ErrMsgUserExitis}
}

//SpawnErrUserFailCap doc
func SpawnErrUserFailCap() message.Response {
	return message.Response{Code: ErrUserFailCap, Data: ErrMsgUserFailCap}
}
