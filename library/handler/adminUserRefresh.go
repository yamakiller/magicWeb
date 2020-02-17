package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserRefresh admin logined refresh token
func AdminUserRefresh(context *gin.Context,
	cacheDB int,
	tokenSecret string,
	tokenExpireMinute int) (string, *protocol.Response) {
	var token string
	var errResult protocol.Response
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		log.Debug("authorization refresh token invalid:%s", err.Error())
		errResult = code.SpawnErrTokenInvalid()
		goto fail
	}

	_, err = database.GetRdsOnlineAdminAccount(cacheDB, tokenUser.ID)
	if err != nil {
		log.Debug("authorization refresh token error:%s", err.Error())
		errResult = code.SpawnErrOnlineUserNot()
		goto fail
	}

	token, err = auth.Enter(tokenSecret,
		tokenUser.ID,
		tokenUser.Account,
		tokenUser.Password,
		tokenExpireMinute)
	if err != nil {
		log.Debug("authorization refresh token error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	if err = database.WithRdsOnlineAdminToken(cacheDB,
		tokenUser.ID,
		token,
		tokenExpireMinute*60*2); err != nil {
		log.Debug("authorization refresh token error:%s", err.Error())
		errResult = code.SpawnErrSystem()
		goto fail
	}

	return token, nil

fail:
	return "", &errResult
}
