package common

import "github.com/gin-gonic/gin"

//GetAdminUserID Returns logined admin user id
func GetAdminUserID(context *gin.Context, db int, tokenSecret string) (string, error) {
	tokenUser, err := GetRequestToken(context, tokenSecret)
	if err != nil {
		return "", err
	}

	return tokenUser.ID, nil
}
