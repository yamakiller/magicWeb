package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminProfileDelete admin user profile delete
func AdminProfileDelete(context *gin.Context,
	sqlHandle string,
	userID string,
	id string) *protocol.Response {

	if err := database.AdminProfileDelete(sqlHandle, id); err != nil {
		logger.Error(0, "admin profile delete error:%s", err.Error)
		errResult := code.SpawnErrDbAbnormal()
		return &errResult
	}

	if err := database.AdminBehaviorAppend(sqlHandle, userID, fmt.Sprintf("删除权限组：%s", id)); err != nil {
		logger.Error(0, "append admin profile behavior log error:%s", err.Error)
	}

	return nil
}
