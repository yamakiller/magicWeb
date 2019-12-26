package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/message"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminBehaviorQueryPage admin behaviro query
func AdminBehaviorQueryPage(context *gin.Context,
	sqlHandle string,
	page,
	pageSize int,
	where string,
	args ...interface{}) ([]models.AdminBehavior, int, *message.Response) {

	behaviors, total, err := database.AdminBehaviorQuery(sqlHandle, page, pageSize, where, args...)
	if err != nil {
		logger.Error(0, "query admin behavior log error:%s", err.Error)
		errResult := code.SpawnErrDbAbnormal()
		return nil, 0, &errResult
	}

	return behaviors, total, nil
}
