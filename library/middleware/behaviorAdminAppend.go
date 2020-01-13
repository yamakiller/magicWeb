package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
)

//BehaviorAdminAppend admin user behavior saver
func BehaviorAdminAppend(context *gin.Context, sqlHandle, tokenSecret, behavior string) {
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		logger.Debug(0, "authorization admin token invalid:%+v", err)
		return
	}

	behavior = fmt.Sprintf("%s,%s", tokenUser.Account, behavior)

	if err := database.AdminBehaviorAppend(sqlHandle, tokenUser.ID, behavior); err != nil {
		logger.Debug(0, "save admin behavior %s error:%+v", behavior, err)
		return
	}
}
