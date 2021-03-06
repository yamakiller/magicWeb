package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
)

//AuthAdminLogin doc Verify login
//Param  (*gin.Context) context
//Param  (int) db
//Param  (string) token secret
//Param  (int)    token timeout
//Param  (bool)   release
func AuthAdminLogin(context *gin.Context,
	db int,
	tokenSecret string,
	release bool) {

	if !release {
		context.Next()
		return
	}

	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		common.ResponseError(context, code.SpawnErrTokenInvalid())
		log.Debug("authorization login token invalid:%+v", err)
		return
	}

	if tokenUser == nil || !tokenUser.IsValid {
		common.ResponseError(context, code.SpawnErrTokenInvalid())
		log.Debug("authorization login token invalid")
		return
	}

	tokenUserAccount, err := database.GetRdsOnlineAdminAccount(db, tokenUser.ID)
	if err != nil {
		common.ResponseError(context, code.SpawnErrOnlineUserNot())
		log.Debug("authorization login token account invalid:%+v", err)
		return
	}

	if tokenUserAccount != tokenUser.Account {
		common.ResponseError(context, code.SpawnErrOnlineUserNot())
		log.Debug("authorization login token account error")
		return
	}

	context.Next()
}
