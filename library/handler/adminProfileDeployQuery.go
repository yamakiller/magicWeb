package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/auth"
)

//AdminProfileDeployQuery admin user profile delete
func AdminProfileDeployQuery(context *gin.Context) *auth.ConfigAdminUserProfileItems {
	return auth.GetAdminProfileDeploy()
}
