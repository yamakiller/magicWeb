package database

import (
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminUserLockOper lock admin user
func AdminUserLockOper(sqlHandle, id string, oper int) error {
	user := models.AdminUser{ID: id, State: uint8(oper)}
	return mysql.Instance().DB(sqlHandle).Model(&user).Where("id = ?", id).Update("state", oper).Error
}
