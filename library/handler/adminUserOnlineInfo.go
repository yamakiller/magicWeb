package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserOnlineInfo Return Online informat
func AdminUserOnlineInfo(context *gin.Context, db int, tokenSecret string) (account string, role []string, res *protocol.Response) {
	var errResult protocol.Response
	var tmpRoles string
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		errResult = code.SpawnErrTokenInvalid()
		goto fail
	}

	account = tokenUser.Account
	tmpRoles, _ = database.GetRdsOnlineAdminRole(db, tokenUser.ID)
	if err = util.JSONUnFormSerialize(tmpRoles, role); err != nil {
		errResult = code.SpawnErrSystemMsg(err.Error())
		goto fail
	}
fail:
	res = &errResult
	return
}
