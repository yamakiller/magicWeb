package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
)

//AuthAdmin Verify if you are a background administrator
//Param  (*gin.Context) context
//Param  (int) db
//Param  (string) token secret
//Param  (bool)   release
func AuthAdmin(context *gin.Context, db int, tokenSecret string, release bool) {
	if !release {
		context.Next()
		return
	}

	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		common.ResponseError(context, code.SpawnErrTokenInvalid())
		log.Debug("authorization admin token invalid:%+v", err)
		return
	}
	backstage, err := database.GetRdsOnlineAdminBackstage(db, tokenUser.ID)
	if err != nil {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		log.Debug("authorization admin get backstate error:%+v", err)
		return
	}

	if backstage < 1 {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		log.Debug("authorization admin get backstate %d", backstage)
		return
	}

	context.Next()
}
