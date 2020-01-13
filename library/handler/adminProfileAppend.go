package handler

import (
	"time"

	"github.com/yamakiller/magicLibs/logger"

	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/protocol"

	"github.com/yamakiller/magicWeb/library/database"

	"github.com/yamakiller/magicLibs/util"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminProfileAppend admin append profile
func AdminProfileAppend(context *gin.Context,
	sqlHandle string,
	userID string,
	name string,
	profiles *auth.AdminUserProfileItems) (string, *protocol.Response) {
	var errResult protocol.Response

	profile := models.AdminProfile{}
	profile.Name = name
	profile.Order = 0
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = profile.CreatedAt
	profile.Data = util.JSONSerialize(profiles)
	nid, err := database.AdminProfileAppend(sqlHandle, &profile)
	if err != nil {
		logger.Error(0, "append admin profile error:%s", err.Error)
		errResult = code.SpawnErrDbAbnormal()
		goto fail
	}

	return nid, nil
fail:
	return "", &errResult
}
