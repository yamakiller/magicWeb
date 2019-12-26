package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yamakiller/magicWeb/library/code"
	"github.com/yamakiller/magicWeb/library/database"
	"github.com/yamakiller/magicWeb/library/message"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminUserQueryPage query admin profile
func AdminUserQueryPage(context *gin.Context,
	sqlHandle string,
	account string,
	order string,
	page,
	pageSize int) ([]models.AdminUser, int, *message.Response) {

	if order == "d" {
		order = "desc"
	} else {
		order = "asc"
	}

	users, total, err := database.AdminUserQuery(sqlHandle, account, order, page, pageSize)
	if err != nil {
		errResult := code.SpawnErrDbAbnormal()
		return nil, 0, &errResult
	}

	return users, total, nil
}
