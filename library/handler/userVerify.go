package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	captchaCtrl "github.com/yamakiller/magicWeb/library/captcha"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/message"
	"github.com/yamakiller/magicWeb/library/models"
)

//VerifyUser verify user login
func VerifyUser(context *gin.Context, account, password, captcha, captchaKey string) *message.Response {
	var errResult message.Response
	var captchaID interface{}
	session := sessions.Default(context)

	if !util.VerifyCaptchaFormat(captcha) {
		logger.Debug(0, "Login Captcha Format error:%s", captcha)
		errResult = code.SpawnErrCaptchaFormat("must consist of eight characters")
		goto fail
	}

	if !util.VerifyAccountFormat(account) &&
		!util.VerifyEmailFormat(account) &&
		!util.VerifyMobileFormat(account) {

		logger.Debug(0, "Login UserName Format error:%s", account)
		errResult = code.SpawnErrUserNameFormat()
		goto fail
	}

	if !util.VerifyPasswordFormat(password) {
		logger.Debug(0, "Login Password Format error:%s", password)
		errResult = code.SpawnErrPwdFormat()
		goto fail
	}

	captchaID = session.Get(captchaKey)
	if captchaID == nil || captchaID == "" {
		logger.Error(0, "Login Captcha not Generate")
		errResult = code.SpawnErrNeedGenerateCaptcha()
		goto fail
	}

	if !captchaCtrl.VerfiyCaptcha(captchaID.(string), captcha) {
		logger.Debug(0, "Login Captcha error:=>%s", captcha)
		errResult = code.SpawnErrCaptcha()
		goto fail
	}
	return nil
fail:
	return &errResult
}

//VerifyUserSignIn verify user informat and Returns token
func VerifyUserSignIn(db int, tokenSecret, account, password string, usr *models.User, loginExpire int) (string, *message.Response) {
	var errResult message.Response
	var token string
	p, err := util.AesEncrypt(usr.Secret, account)
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

	if err = database.CreateRdsOnlineUserVal(db,
		usr.ID,
		token,
		usr.Account,
		usr.Password,
		usr.Secret,
		usr.Profile.Data,
		util.TimeNowFormat(),
		int(usr.Backstage),
		loginExpire); err != nil {
		logger.Error(0, "Login auth redis error:%+v", err)
		errResult = code.SpawnErrSystem()
		goto fail
	}

	return token, nil
fail:
	return "", &errResult
}
