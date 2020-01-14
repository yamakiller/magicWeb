package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
)

//GetRequestTokenRole Returns online User role
func GetRequestTokenRole(context *gin.Context, db int, tokenSecret string) ([]string, error) {
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		return nil, err
	}

	role, err := database.GetRdsOnlineAdminRole(db, tokenUser.ID)
	if err != nil {
		return nil, err
	}

	var roles []string
	if err = util.JSONUnFormSerialize(role, roles); err != nil {
		return nil, err
	}

	return roles, nil
}
