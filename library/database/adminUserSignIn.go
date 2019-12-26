package database

import (
	"github.com/jinzhu/gorm"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//GetAdminUserSignIn Return admin user
func GetAdminUserSignIn(account string, sqlHandle string) (*models.AdminUser, error) {
	result := models.AdminUser{}
	if err := mysql.Instance().DB(sqlHandle).Where("account=?", account).First(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
