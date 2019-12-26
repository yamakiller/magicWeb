package database

import (
	"github.com/jinzhu/gorm"

	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AlreadyAdminAccount is exitis check account or email, or mobile
func AlreadyAdminAccount(sqlHandle string, account string) (bool, error) {
	result := models.AdminUser{}
	if err := mysql.Instance().
		DB(sqlHandle).
		Select("id").
		Where("account=?", account).
		Or("email=?", account).Or("mobile=?", account).First(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}

		return false, err
	}
	return true, nil
}

//AlreadOnlyAdminAccount is exitis check account
func AlreadOnlyAdminAccount(sqlHandle string, account string) (bool, error) {
	result := models.AdminUser{}
	if err := mysql.Instance().
		DB(sqlHandle).
		Select("id").
		Where("account=?", account).
		First(&result).Error; err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}

		return false, err
	}
	return true, nil
}

//CreateAdminAccount create account
func CreateAdminAccount(sqlHandle string, user *models.AdminUser) error {
	return mysql.Instance().DB(sqlHandle).Create(user).Error
}
