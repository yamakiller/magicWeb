package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
)

//AuthAdminAccess Verify that you have access
//Param  (*gin.Context) context
func AuthAdminAccess(context *gin.Context, db int, tokenSecret string, release bool) {
	if !release {
		context.Next()
		return
	}

	profileItems, err := GetRequestTokenProfile(context, db, tokenSecret)
	if err != nil {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		logger.Debug(0, "authorization profile access error:%+v", err)
		return
	}

	if !auth.VerifyAdminProfile(profileItems.Items, context.Request.RequestURI) {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		logger.Debug(0, "authorization profile access %s need access", context.Request.RequestURI)
		return
	}

	context.Next()
}
