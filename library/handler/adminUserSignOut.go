package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/message"
)

//AdminSignOut admin user sign-out
func AdminSignOut(context *gin.Context,
	cacheDB int,
	tokenSecret string) *message.Response {

	var errResult message.Response
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
