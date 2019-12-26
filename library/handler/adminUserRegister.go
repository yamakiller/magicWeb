package handler

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/message"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminUserRegister admin user register
func AdminUserRegister(context *gin.Context,
	sqlHandle,
	captchaKey,
	account,
	pwd,
	againPwd,
	profileID,
	captchaVal string) *message.Response {
	var errResult message.Response
	var err error
	var is bool
	var secret, password string
	var rAdminUsr models.AdminUser
	var captchaID interface{}

	if !util.VerifyCaptchaFormat(captchaVal) {
		logger.Debug(0, "Regiser account captcha format error:%s", captchaVal)
		errResult = code.SpawnErrCaptchaFormat("must consist of eight characters")
		goto fail
	}

	if !util.VerifyAccountFormat(account) {
		logger.Debug(0, "Register account format error:%s", account)
		errResult = code.SpawnErrUserNameFormat()
		goto fail
	}

	if !util.VerifyPasswordFormat(pwd) {
		logger.Debug(0, "Register password format error:%s", pwd)
		errResult = code.SpawnErrPwdFormat()
		goto fail
	}

	if pwd != againPwd {
		logger.Debug(0, "Register password first!=second")
		errResult = code.SpawnErrPwdAgin()
		goto fail
	}

	captchaID = sessions.Default(context).Get(captchaKey)
	if captchaID == nil || captchaID == "" {
		logger.Error(0, "SignIn Admin Captcha not Generate")
		errResult = code.SpawnErrNeedGenerateCaptcha()
		goto fail
	}

	if is, err = database.AlreadOnlyAdminAccount(sqlHandle, account); is || err != nil {
		if is {
			errResult = code.SpawnErrUserExitis()
			goto fail
		}
		logger.Debug(0, "Register account error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	secret = util.RandStr(8)
	password, err = util.AesEncrypt(secret, pwd)
	if err != nil {
		logger.Error(0, "Register password encrypt error:%s", err.Error)
		errResult = code.SpawnErrSystem()
		goto fail
	}

	rAdminUsr.ID = util.SpawnUUID()
	rAdminUsr.Account = account
	rAdminUsr.Password = password
	rAdminUsr.Nick = account
	rAdminUsr.Secret = secret
	rAdminUsr.State = 0
	rAdminUsr.Backstage = 1
	rAdminUsr.ProfileID = profileID
	rAdminUsr.Source = "local"
	rAdminUsr.CreatedAt = time.Now()
	rAdminUsr.FailLastTime = time.Time{}
	rAdminUsr.LoginedLastTime = rAdminUsr.CreatedAt
	rAdminUsr.LoginedIP = context.ClientIP()
	rAdminUsr.CreateIP = rAdminUsr.LoginedIP

	if err := database.CreateAdminAccount(sqlHandle, &rAdminUsr); err != nil {
		logger.Error(0, "Register database error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	return nil
fail:
	return &errResult
}
