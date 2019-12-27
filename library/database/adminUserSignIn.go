package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//GetAdminUserSignIn Return admin user
func GetAdminUserSignIn(account, sqlHandle string) (*models.AdminUser, error) {
	result := models.AdminUser{}
	if err := mysql.Instance().DB(sqlHandle).Where("account=?", account).First(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

//WithAdminUserSignInSuccess Update signIn success state
func WithAdminUserSignInSuccess(account, ip, sqlHandle string) error {
	err := mysql.Instance().DB(sqlHandle).Model(models.AdminUser{}).
		Where("account=?", account).
		Update(map[string]interface{}{"logined": gorm.Expr("logined+?", 1), "logined_last_time": time.Now(), "logined_ip": ip, "fail": 0}).Error
	return err
}

//WithAdminUserSignInFail Update signIn password error state
func WithAdminUserSignInFail(account, sqlHandle string) error {
	err := mysql.Instance().DB(sqlHandle).Model(models.AdminUser{}).
		Where("account=?", account).
		Update(map[string]interface{}{"fail": gorm.Expr("fail+?", 1), "fail_last_time": time.Now()}).Error
	return err
}
