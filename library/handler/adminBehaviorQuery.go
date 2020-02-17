package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/log"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminBehaviorQueryPage admin behaviro query
func AdminBehaviorQueryPage(context *gin.Context,
	sqlHandle string,
	page,
	pageSize int,
	where string,
	args ...interface{}) ([]models.AdminBehavior, int, *protocol.Response) {

	behaviors, total, err := database.AdminBehaviorQuery(sqlHandle, page, pageSize, where, args...)
	if err != nil {
		log.Error("query admin behavior log error:%s", err.Error)
		errResult := code.SpawnErrDbAbnormal()
		return nil, 0, &errResult
	}

	return behaviors, total, nil
}
