package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserLock admin user lock
func AdminUserLock(context *gin.Context,
	cacheDB int,
	sqlHandle string,
	adminUserID string,
	adminUserState int) *protocol.Response {

	var isOnline bool
	var newState int
	var errResult protocol.Response

	oldState, err := database.AdminUserQueryState(sqlHandle, adminUserID)
	if err != nil {
		errResult = code.SpawnErrSystemMsg(err.Error())
		goto fail
	}

	if _, err = database.GetRdsOnlineAdminAccount(cacheDB, adminUserID); err != nil {
		if err != database.ErrOnlineUserEmpty {
			errResult = code.SpawnErrSystemMsg(err.Error())
			goto fail
		}
	} else {
		isOnline = true
	}

	if oldState == adminUserState {
		return nil
	}

	if err = database.AdminUserLockOper(sqlHandle, adminUserID, adminUserState); err != nil {
		errResult = code.SpawnErrSystemMsg(err.Error())
		goto fail
	}

	if isOnline && newState != 0 {
		if err != database.RemoveOnlineAdminUser(cacheDB, adminUserID) {
			log.Debug("remove online admn error:%s", err.Error())
		}
	}

	return nil
fail:
	return &errResult
}
