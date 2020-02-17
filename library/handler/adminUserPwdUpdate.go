package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/encryption/aes"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserPwdUpdate admin user modify password
func AdminUserPwdUpdate(context *gin.Context,
	cacheDB int,
	sqlHandle,
	adminUserID,
	adminUserPwd,
	adminUserAginPwd string) *protocol.Response {

	var pwd, oldPwd string
	var err error
	var oldUser *models.AdminUser
	var errResult protocol.Response

	if !util.VerifyPasswordFormat(pwd) {
		log.Debug("Modify admin user password format error:%s", pwd)
		errResult = code.SpawnErrPwdFormat()
		goto fail
	}

	if pwd != adminUserAginPwd {
		log.Debug("Modify admin user password first!=second")
		errResult = code.SpawnErrPwdAgin()
		goto fail
	}

	oldUser, err = database.AdminUserQueryPwd(sqlHandle, adminUserID)
	if err != nil {
		errResult = code.SpawnErrSystemMsg(err.Error())
		goto fail
	}

	pwd, err = aes.Encrypt(oldUser.Secret, adminUserPwd)

	if err = database.AdminUserPwdUpdate(sqlHandle, adminUserID, pwd); err != nil {
		errResult = code.SpawnErrSystemMsg(err.Error())
		goto fail
	}

	if oldPwd, err = database.GetRdsOnlineAdminPassword(cacheDB, adminUserID); err == nil {
		if oldPwd != pwd {
			if err = database.WithRdsOnlineAdminPwd(cacheDB, adminUserID, pwd); err != nil {
				log.Debug("Modify admin user online password error:%s", err.Error())
			}
		}
	}

	return nil
fail:
	return &errResult
}
