package code

const (
	//ErrParam param error of code
	ErrParam = int(-1)
	//ErrDbAbnormal db/mysql abnormal error of code
	ErrDbAbnormal = int(-2)
	//ErrSystem  system error of code
	ErrSystem = int(-3)
	//ErrCaptchaFormat captcha format error of code
	ErrCaptchaFormat = int(-4)
	//ErrUserNameFormat username format error of code
	ErrUserNameFormat = int(-5)
	//ErrPwdFormat password format error of code
	ErrPwdFormat = int(-6)
	//ErrUserName username error of code
	ErrUserName = int(-7)
	//ErrPwd password error of code
	ErrPwd = int(-8)
	//ErrPwdAgain password again error of code
	ErrPwdAgain = int(-9)
	//ErrTokenNot authorization token error of code
	ErrTokenNot = int(-10)
	//ErrTokenInvalid authorization token invalid error of code
	ErrTokenInvalid = int(-11)
	//ErrOnlineUserNot user online data not found error of code
	ErrOnlineUserNot = int(-12)
	//ErrNeedPerm user need sufficient permissions error of code
	ErrNeedPerm = int(-13)
	//ErrNeedGenerateCaptcha captcha not generate error of code
	ErrNeedGenerateCaptcha = int(-14)
	//ErrCaptcha captcha error of code
	ErrCaptcha = int(-15)
	//ErrConfigNot config data non-existent error of code
	ErrConfigNot = int(-16)
	//ErrUserExitis user already exists error of code
	ErrUserExitis = int(-17)
)

const (
	//ErrMsgParam param error of message
	ErrMsgParam = "param error"
	//ErrMsgSystem system error of message
	ErrMsgSystem = "system error"
	//ErrMsgDbAbnormal db/mysql abnormal error of message
	ErrMsgDbAbnormal = "system abnormal"
	//ErrMsgCaptchaFormat captcha format error of message
	ErrMsgCaptchaFormat = "captcha format error"
	//ErrMsgUserNameFormat username format error of message
	ErrMsgUserNameFormat = "username format error"
	//ErrMsgPwdFormat password format error of message
	ErrMsgPwdFormat = "password format error"
	//ErrMsgUserName username error of message
	ErrMsgUserName = "user name error"
	//ErrMsgPwd password error of message
	ErrMsgPwd = "password error"
	//ErrMsgPwdAgain password again error of message
	ErrMsgPwdAgain = "passwords entered twice do not match"
	//ErrMsgTokenNot authorization token error of message
	ErrMsgTokenNot = "authorization token does not exist"
	//ErrMsgTokenInvalid authorization token invalid error of message
	ErrMsgTokenInvalid = "authorization token invalid"
	//ErrMsgOnlineUserNot user online data not found error of message
	ErrMsgOnlineUserNot = "user is not logged in"
	//ErrMsgNeedPerm user need sufficient permissions error of message
	ErrMsgNeedPerm = "need sufficient permissions"
	//ErrMsgNeedGenerateCaptcha captcha not generate error of message
	ErrMsgNeedGenerateCaptcha = "need generate captcha"
	//ErrMsgCaptcha captcha error of message
	ErrMsgCaptcha = "captcha is error"
	//ErrMsgConfigNot config data non-existent error of message
	ErrMsgConfigNot = "config data non-existent"
	//ErrMsgUserExitis user already exists error of message
	ErrMsgUserExitis = "user already exists"
)
