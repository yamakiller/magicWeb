package database

import (
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminUserPwdUpdate lock admin user
func AdminUserPwdUpdate(sqlHandle, id string, pwd string) error {
	user := models.AdminUser{ID: id, Password: pwd}
	return mysql.Instance().DB(sqlHandle).Model(&user).Where("id = ?", id).Update("password", pwd).Error
}
