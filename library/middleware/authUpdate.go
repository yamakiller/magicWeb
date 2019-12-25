package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
)

//AuthUpdate Verify that you have update
//Param  (*gin.Context) context
func AuthUpdate(context *gin.Context, db int, tokenSecret string, release bool) {
	if !release {
		context.Next()
		return
	}

	profileItems, err := GetRequestTokenProfile(context, db, tokenSecret)
	if err != nil {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		logger.Debug(0, "authorization profile update error:%+v", err)
		return
	}

	if !auth.VerifyProfile(profileItems.Items, context.Request.RequestURI, auth.ProfileUpdate) {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		logger.Debug(0, "authorization profile update %s need update", context.Request.RequestURI)
		return
	}

	context.Next()
}
