package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserSignOut admin user sign-out
func AdminUserSignOut(context *gin.Context,
	cacheDB int,
	tokenSecret string) *protocol.Response {

	var errResult protocol.Response
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		logger.Debug(0, "authorization sign out token invalid:%s", err.Error())
		errResult = code.SpawnErrTokenInvalid()
		goto fail
	}

	if err := database.RemoveOnlineAdminUser(cacheDB, tokenUser.ID); err != nil {
		logger.Debug(0, "authorization sign out error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	return nil
fail:
	return &errResult
}
