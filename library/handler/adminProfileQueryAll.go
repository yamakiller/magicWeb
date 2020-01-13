package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminProfileQueryAll query admin profile group all
func AdminProfileQueryAll(context *gin.Context,
	sqlHandle string) ([]models.AdminProfile, *protocol.Response) {
	var errResult protocol.Response
	profiles, err := database.AdminProfileQueryAll(sqlHandle)
	if err != nil {
		logger.Error(0, "all query admin profile error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	return profiles, nil
fail:
	return nil, &errResult
}
