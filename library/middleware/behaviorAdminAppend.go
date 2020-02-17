package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
)

//BehaviorAdminAppend admin user behavior saver
func BehaviorAdminAppend(context *gin.Context, sqlHandle, tokenSecret, behavior string) {
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		log.Debug("authorization admin token invalid:%+v", err)
		return
	}

	behavior = fmt.Sprintf("%s,%s", tokenUser.Account, behavior)

	if err := database.AdminBehaviorAppend(sqlHandle, tokenUser.ID, behavior); err != nil {
		log.Debug("save admin behavior %s error:%+v", behavior, err)
		return
	}
}
