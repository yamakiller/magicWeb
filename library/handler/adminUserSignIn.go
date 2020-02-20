package handler

import (
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/encryption/aes"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/captcha"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserSignIn admin user sign-in
func AdminUserSignIn(context *gin.Context,
	cacheDB int,
	sqlHandle,
	tokenSecret,
	captchaKey string,
	tokenExpireMinute int,
	failCap int,
	failExpire int,
	account,
	password,
	captchaVal string) (*protocol.Response, string) {
	var errResult protocol.Response
	var captchaID interface{}
	var usr *models.AdminUser
	var err error
	var p, token string

	if !util.VerifyCaptchaFormat(captchaVal) {
		log.Debug("SignIn Admin Captcha Format error:%s", captchaVal)
		errResult = code.SpawnErrCaptchaFormat("must consist of eight characters")
		goto fail
	}

	if !util.VerifyAccountFormat(account) &&
		!util.VerifyEmailFormat(account) &&
		!util.VerifyMobileFormat(account) {

		log.Debug("SignIn Admin UserName Format error:%s", account)
		errResult = code.SpawnErrUserNameFormat()
		goto fail
	}

	if !util.VerifyPasswordFormat(password) {
		log.Debug("SignIn Admin Password Format error:%s", password)
		errResult = code.SpawnErrPwdFormat()
		goto fail
	}

	captchaID = sessions.Default(context).Get(captchaKey)
	if captchaID == nil || captchaID == "" {
		log.Error("SignIn Admin Captcha not Generate")
		errResult = code.SpawnErrNeedGenerateCaptcha()
		goto fail
	}

	if !captcha.VerfiyCaptcha(captchaID.(string), captchaVal) {
		log.Debug("SignIn Admin Captcha error:=>%s", captchaVal)
		errResult = code.SpawnErrCaptcha()
		goto fail
	}

	usr, err = database.GetAdminUserSignIn(account, sqlHandle)
	if err != nil {
		log.Error("SignIn Admin fail:%s", err.Error())
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	if usr == nil {
		log.Debug("SignIn Admin user does not exist:%s", account)
		errResult = code.SpawnErrUserName()
		goto fail
	}

	if int(usr.Fail) > failCap && (time.Now().Unix()-usr.FailLastTime.Unix()) < int64(failExpire) {
		log.Debug("SignIn Admin fail limit:[%d:%d]",
			usr.Fail,
			(time.Now().Unix() - usr.FailLastTime.Unix()))
		errResult = code.SpawnErrUserFailCap()
		goto fail
	}

	p, err = aes.Encrypt(usr.Secret, account)
	if err != nil {
		log.Error("SignIn Admin fail encrypt error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if p != usr.Password {
		log.Debug("SignIn Admin password error:%s", account)
		errResult = code.SpawnErrPwd()
		if err = database.WithAdminUserSignInFail(usr.Account, sqlHandle); err != nil {
			log.Error("SignIn Admin fail update state error:%s", err.Error())
		}
		goto fail
	}

	token, err = auth.Enter(tokenSecret,
		usr.ID,
		usr.Account,
		usr.Password,
		tokenExpireMinute)

	if err != nil {
		log.Error("SignIn Admin jwt error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if err = database.CreateRdsOnlineAdminUserVal(cacheDB,
		usr.ID,
		token,
		usr.Account,
		usr.Password,
		usr.Secret,
		usr.Roles,
		util.TimestampFormat(),
		int(usr.Backstage),
		tokenExpireMinute*60*2); err != nil {
		log.Error("SignIn Admin auth redis error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if err = database.WithAdminUserSignInSuccess(usr.Account, context.ClientIP(), sqlHandle); err != nil {
		log.Error("SignIn Admin Complate update state error:%s", err.Error())
	}

	return nil, token
fail:
	return &errResult, ""
}
