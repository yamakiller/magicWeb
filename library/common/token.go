package common

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/auth"
)

//GetRequestToken Returns *Claims
func GetRequestToken(context *gin.Context, tokenSecret string) (*auth.Claims, error) {
	token := context.Request.Header.Get("token")
	if token == "" {
		return nil, errors.New("token unknow")
	}
	tokenUser, err := auth.Get(tokenSecret, token)
	if err != nil {
		return nil, err
	}

	return tokenUser, nil
}
