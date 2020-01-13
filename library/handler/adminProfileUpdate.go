package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/models"
	"github.com/yamakiller/magicWeb/library/protocol"
)

//AdminProfileUpdate admin update profile
func AdminProfileUpdate(context *gin.Context,
	sqlHandle string,
	userID string,
	id string,
	name string,
	profiles *auth.AdminUserProfileItems) *protocol.Response {
	var errResult protocol.Response

	profile := models.AdminProfile{}
	profile.ID = id
	profile.Name = name
	profile.Order = 0
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = profile.CreatedAt
	profile.Data = util.JSONSerialize(profiles)
	err := database.AdminProfileUpdate(sqlHandle, &profile)
	if err != nil {
		logger.Error(0, "update admin profile error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	return nil
fail:
	return &errResult
}
