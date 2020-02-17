package handler

import (
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/yamakiller/magicLibs/encryption/aes"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserRegister admin user register
func AdminUserRegister(context *gin.Context,
	sqlHandle,
	captchaKey,
	account,
	pwd,
	againPwd,
	roles,
	Identity,
	captchaVal string) *protocol.Response {
	var errResult protocol.Response
	var err error
	var is bool
	var secret, password string
	var rAdminUsr models.AdminUser
	var captchaID interface{}

	if !util.VerifyCaptchaFormat(captchaVal) {
		log.Debug("Regiser account captcha format error:%s", captchaVal)
		errResult = code.SpawnErrCaptchaFormat("must consist of eight characters")
		goto fail
	}

	if !util.VerifyAccountFormat(account) {
		log.Debug("Register account format error:%s", account)
		errResult = code.SpawnErrUserNameFormat()
		goto fail
	}

	if !util.VerifyPasswordFormat(pwd) {
		log.Debug("Register password format error:%s", pwd)
		errResult = code.SpawnErrPwdFormat()
		goto fail
	}

	if pwd != againPwd {
		log.Debug("Register password first!=second")
		errResult = code.SpawnErrPwdAgin()
		goto fail
	}

	captchaID = sessions.Default(context).Get(captchaKey)
	if captchaID == nil || captchaID == "" {
		log.Error("SignIn Admin Captcha not Generate")
		errResult = code.SpawnErrNeedGenerateCaptcha()
		goto fail
	}

	if is, err = database.AlreadOnlyAdminAccount(sqlHandle, account); is || err != nil {
		if is {
			errResult = code.SpawnErrUserExitis()
			goto fail
		}
		log.Debug("Register account error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	secret = util.RandStr(16)
	password, err = aes.Encrypt(secret, pwd)
	if err != nil {
		log.Error("Register password encrypt error:%s", err.Error)
		errResult = code.SpawnErrSystem()
		goto fail
	}

	rAdminUsr.ID = util.SpawnUUID()
	rAdminUsr.Account = account
	rAdminUsr.Password = password
	rAdminUsr.Nick = account
	rAdminUsr.Identity = Identity
	rAdminUsr.Secret = secret
	rAdminUsr.State = 0
	rAdminUsr.Backstage = 1
	rAdminUsr.Roles = roles
	rAdminUsr.Source = "local"
	rAdminUsr.CreatedAt = time.Now()
	rAdminUsr.FailLastTime = time.Time{}
	rAdminUsr.LoginedLastTime = rAdminUsr.CreatedAt
	rAdminUsr.LoginedIP = context.ClientIP()
	rAdminUsr.CreateIP = rAdminUsr.LoginedIP

	if err := database.CreateAdminAccount(sqlHandle, &rAdminUsr); err != nil {
		log.Error("Register database error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	return nil
fail:
	return &errResult
}
