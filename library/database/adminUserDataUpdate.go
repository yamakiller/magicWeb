package database

import (
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminUserDataUpdate update admin user data
func AdminUserDataUpdate(sqlHandle string, user *models.AdminUser) error {
	return mysql.Instance().DB(sqlHandle).Model(user).Updates(models.AdminUser{ID: user.ID,
		Nick:      user.Nick,
		Email:     user.Email,
		Mobile:    user.Mobile,
		Identity:  user.Identity,
		ProfileID: user.ProfileID}).Error
}
