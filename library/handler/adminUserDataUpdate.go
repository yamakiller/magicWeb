package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminUserDataUpdate admin user modify data
func AdminUserDataUpdate(context *gin.Context,
	sqlHandle string,
	adminUser *models.AdminUser) *protocol.Response {

	var errResult protocol.Response
	if err := database.AdminUserDataUpdate(sqlHandle, adminUser); err != nil {
		errResult = code.SpawnErrSystemMsg(err.Error)
		goto fail
	}

	return nil
fail:
	return &errResult
}
