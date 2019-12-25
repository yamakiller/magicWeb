package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/database"
)

//GetRequestTokenProfile Returns online User profile
func GetRequestTokenProfile(context *gin.Context, db int, tokenSecret string) (*auth.UserProfileItems, error) {
	tokenUser, err := common.GetRequestToken(context, tokenSecret)
	if err != nil {
		return nil, err
	}

	profile, err := database.GetRdsOnlineProfile(db, tokenUser.ID)
	if err != nil {
		return nil, err
	}

	result := auth.UserProfileItems{}
	if err := util.JSONUnFormSerialize(profile, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
