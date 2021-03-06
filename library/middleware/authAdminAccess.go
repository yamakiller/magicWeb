package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/log"
)

//AuthAdminAccess Verify that you have access
//Param  (*gin.Context) context
func AuthAdminAccess(context *gin.Context, db int, tokenSecret string, needs []string, release bool) {
	if !release {
		context.Next()
		return
	}

	roles, err := GetRequestTokenRole(context, db, tokenSecret)
	if err != nil {
		common.ResponseError(context, code.SpawnErrNeedPerm())
		log.Debug("authorization role access error:%+v", err)
		return
	}

	for _, need := range needs {
		for _, role := range roles {
			if strings.ToLower(need) == strings.ToLower(role) {
				goto next
			}
		}
	}

	common.ResponseError(context, code.SpawnErrNeedPerm())
	log.Debug("authorization profile access %s need access",
		context.Request.RequestURI)
	return
next:
	context.Next()
}
