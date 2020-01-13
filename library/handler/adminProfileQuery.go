package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminProfileQueryPage query admin profile
func AdminProfileQueryPage(context *gin.Context,
	sqlHandle string,
	page,
	pageSize int) ([]models.AdminProfile, int, *protocol.Response) {
	var errResult protocol.Response
	profiles, total, err := database.AdminProfileQuery(sqlHandle, page, pageSize)
	if err != nil {
		logger.Error(0, "page query admin profile error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	return profiles, total, nil
fail:
	return nil, 0, &errResult
}

//AdminProfileQuery
