package code

import (
	"fmt"

	"github.com/yamakiller/magicWeb/library/protocol"
)

//SpawnResponse doc
//ummary create success  response
func SpawnResponse(data interface{}) protocol.Response {
	return protocol.Response{Code: 20000, Data: data}
}

//SpawnErrParam doc
//Summary create param error response json
//Return (protocol.Response)
func SpawnErrParam() protocol.Response {
	return protocol.Response{Code: ErrParam, Message: ErrMsgParam}
}

//SpawnErrSystem doc
//Summary create system error response json
//Return (protocol.Response)
func SpawnErrSystem() protocol.Response {
	return protocol.Response{Code: ErrSystem, Message: ErrMsgSystem}
}

//SpawnErrSystemMsg doc
//Summary create system error message response json
//Return (protocol.Response)
func SpawnErrSystemMsg(msg string) protocol.Response {
	return protocol.Response{Code: ErrSystem, Message: msg}
}

//SpawnErrDbAbnormal doc
//Summary create db/mysql abnormal error response json
//Return (protocol.Response)
func SpawnErrDbAbnormal() protocol.Response {
	return protocol.Response{Code: ErrDbAbnormal, Message: ErrMsgDbAbnormal}
}

//SpawnErrCaptchaFormat doc
//Summary create captcha format error response json
//Return (protocol.Response)
func SpawnErrCaptchaFormat(msg string) protocol.Response {
	return protocol.Response{Code: ErrCaptchaFormat, Message: fmt.Sprintf("%s:%s", ErrMsgCaptchaFormat, msg)}
}

//SpawnErrUserNameFormat doc
//Summary create username format error response json
//Return (protocol.Response)
func SpawnErrUserNameFormat() protocol.Response {
	return protocol.Response{Code: ErrUserNameFormat, Message: ErrMsgUserNameFormat}
}

//SpawnErrPwdFormat doc
//Summary create password format error response json
//Return (protocol.Response)
func SpawnErrPwdFormat() protocol.Response {
	return protocol.Response{Code: ErrPwdFormat, Message: ErrMsgPwdFormat}
}

//SpawnErrUserName doc
//Summary create user name error response json
//Return (protocol.Response)
func SpawnErrUserName() protocol.Response {
	return protocol.Response{Code: ErrUserName, Message: ErrMsgUserName}
}

//SpawnErrPwd doc
//Summary create password error response json
//Return (protocol.Response)
func SpawnErrPwd() protocol.Response {
	return protocol.Response{Code: ErrPwd, Message: ErrMsgPwd}
}

//SpawnErrPwdAgin doc
//Method SpawnErrPwdAgin @Summary create password again error response json
//Return (protocol.Response)
func SpawnErrPwdAgin() protocol.Response {
	return protocol.Response{Code: ErrPwdAgain, Message: ErrMsgPwdAgain}
}

//SpawnErrTokenNot doc
//Method SpawnErrTokenNot @Summary create authorization token error response json
//Return (protocol.Response)
func SpawnErrTokenNot() protocol.Response {
	return protocol.Response{Code: ErrTokenNot, Message: ErrMsgTokenNot}
}

//SpawnErrTokenInvalid doc
//Method SpawnErrTokenInvalid @Summary create authorization token invalid error response json
//Return (protocol.Response)
func SpawnErrTokenInvalid() protocol.Response {
	return protocol.Response{Code: ErrTokenInvalid, Message: ErrMsgTokenInvalid}
}

//SpawnErrOnlineUserNot doc
//Method SpawnErrOnlineUserNot @Summary create user unonline error response json
//Return (protocol.Response)
func SpawnErrOnlineUserNot() protocol.Response {
	return protocol.Response{Code: ErrOnlineUserNot, Message: ErrMsgOnlineUserNot}
}

//SpawnErrNeedPerm doc
//Method SpawnErrNeedPerm @Summary create need sufficient permissions error  response json
//Return (protocol.Response)
func SpawnErrNeedPerm() protocol.Response {
	return protocol.Response{Code: ErrNeedPerm, Message: ErrMsgNeedPerm}
}

//SpawnErrNeedGenerateCaptcha doc
//Method SpawnErrNeedGenerateCaptcha @Summary create captcha not grenerate error  response json
//Return (protocol.Response)
func SpawnErrNeedGenerateCaptcha() protocol.Response {
	return protocol.Response{Code: ErrNeedGenerateCaptcha, Message: ErrMsgNeedGenerateCaptcha}
}

//SpawnErrCaptcha doc
//Summary create captcha error  response json
//Method SpawnErrCaptcha
//Return (protocol.Response)
func SpawnErrCaptcha() protocol.Response {
	return protocol.Response{Code: ErrCaptcha, Message: ErrMsgCaptcha}
}

//SpawnErrConfigNot doc
//Summary create config data non-existent error response json
//Method SpawnErrConfigNot
//Return (protocol.Response)
func SpawnErrConfigNot() protocol.Response {
	return protocol.Response{Code: ErrConfigNot, Message: ErrMsgConfigNot}
}

//SpawnErrUserExitis doc
func SpawnErrUserExitis() protocol.Response {
	return protocol.Response{Code: ErrUserExitis, Message: ErrMsgUserExitis}
}

//SpawnErrUserFailCap doc
func SpawnErrUserFailCap() protocol.Response {
	return protocol.Response{Code: ErrUserFailCap, Message: ErrMsgUserFailCap}
}
