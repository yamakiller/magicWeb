package handler

import (
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/captcha"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserSignIn admin user sign-in
func AdminUserSignIn(context *gin.Context,
	cacheDB int,
	sqlHandle,
	tokenSecret,
	captchaKey string,
	tokenExpire int,
	failCap int,
	failExpire int,
	account,
	password,
	captchaVal string) (*protocol.Response, *models.AdminUser, string) {
	var errResult protocol.Response
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
		logger.Error(0, "SignIn Admin fail:%s", err.Error())
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	if usr == nil {
		logger.Debug(0, "SignIn Admin user does not exist:%s", account)
		errResult = code.SpawnErrUserName()
		goto fail
	}

	if int(usr.Fail) > failCap && (time.Now().Unix()-usr.FailLastTime.Unix()) < int64(failExpire) {
		logger.Debug(0, "SignIn Admin fail limit:[%d:%d]", usr.Fail, (time.Now().Unix() - usr.FailLastTime.Unix()))
		errResult = code.SpawnErrUserFailCap()
		goto fail
	}

	p, err = util.AesEncrypt(usr.Secret, account)
	if err != nil {
		logger.Error(0, "SignIn Admin fail encrypt error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if p != usr.Password {
		logger.Debug(0, "SignIn Admin password error:%s", account)
		errResult = code.SpawnErrPwd()
		if err = database.WithAdminUserSignInFail(usr.Account, sqlHandle); err != nil {
			logger.Error(0, "SignIn Admin fail update state error:%s", err.Error())
		}
		goto fail
	}

	token, err = auth.Enter(tokenSecret, usr.ID, usr.Account, usr.Password, tokenExpire)
	if err != nil {
		logger.Error(0, "SignIn Admin jwt error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if err = database.CreateRdsOnlineAdminUserVal(cacheDB,
		usr.ID,
		token,
		usr.Account,
		usr.Password,
		usr.Secret,
		usr.Profile.Data,
		util.TimeNowFormat(),
		int(usr.Backstage),
		tokenExpire*2); err != nil {
		logger.Error(0, "SignIn Admin auth redis error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if err = database.WithAdminUserSignInSuccess(usr.Account, context.ClientIP(), sqlHandle); err != nil {
		logger.Error(0, "SignIn Admin Complate update state error:%s", err.Error())
	}

	return nil, usr, token
fail:
	return &errResult, nil, ""
}
