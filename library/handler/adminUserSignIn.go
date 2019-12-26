package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/captcha"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/message"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminSignIn admin user sign-in
func AdminSignIn(context *gin.Context,
	CacheDB int,
	sqlHandle,
	tokenSecret,
	captchaKey string,
	loginExpire int,
	account,
	password,
	captchaVal string) (*message.Response, *models.AdminUser, string) {
	var errResult message.Response
	var captchaID interface{}
	var usr *models.AdminUser
	var err error
	var p, token string

	if !util.VerifyCaptchaFormat(captchaVal) {
		logger.Debug(0, "SignIn Admin Captcha Format error:%s", captchaVal)
		errResult = code.SpawnErrCaptchaFormat("must consist of eight characters")
		goto fail
	}

	if !util.VerifyAccountFormat(account) &&
		!util.VerifyEmailFormat(account) &&
		!util.VerifyMobileFormat(account) {

		logger.Debug(0, "SignIn Admin UserName Format error:%s", account)
		errResult = code.SpawnErrUserNameFormat()
		goto fail
	}

	if !util.VerifyPasswordFormat(password) {
		logger.Debug(0, "SignIn Admin Password Format error:%s", password)
		errResult = code.SpawnErrPwdFormat()
		goto fail
	}

	captchaID = sessions.Default(context).Get(captchaKey)
	if captchaID == nil || captchaID == "" {
		logger.Error(0, "SignIn Admin Captcha not Generate")
		errResult = code.SpawnErrNeedGenerateCaptcha()
		goto fail
	}

	if !captcha.VerfiyCaptcha(captchaID.(string), captchaVal) {
		logger.Debug(0, "SignIn Admin Captcha error:=>%s", captchaVal)
		errResult = code.SpawnErrCaptcha()
		goto fail
	}

	usr, err = database.GetAdminUserSignIn(account, sqlHandle)
	if err != nil {
		logger.Error(0, "SignIn Admin fail:%+v", err)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	if usr == nil {
		logger.Debug(0, "SignIn Admin user does not exist:%s", account)
		errResult = code.SpawnErrUserName()
		goto fail
	}

	p, err = util.AesEncrypt(usr.Secret, account)
	if err != nil {
		logger.Error(0, "Login fail encrypt error:%+v", err)
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if p != usr.Password {
		logger.Debug(0, "Login password error:%s", account)
		errResult = code.SpawnErrPwd()
		goto fail
	}

	token, err = auth.Enter(tokenSecret, usr.ID, usr.Account, usr.Password, loginExpire)
	if err != nil {
		logger.Error(0, "Login auth jwt error:%+v", err)
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if err = database.CreateRdsOnlineAdminUserVal(CacheDB,
		usr.ID,
		token,
		usr.Account,
		usr.Password,
		usr.Secret,
		usr.Profile.Data,
		util.TimeNowFormat(),
		int(usr.Backstage),
		loginExpire*2); err != nil {
		logger.Error(0, "Login auth redis error:%+v", err)
		errResult = code.SpawnErrSystem()
		goto fail
	}

	return nil, usr, token
fail:
	return &errResult, nil, ""
}
