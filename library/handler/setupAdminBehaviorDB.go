package handler

import (
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//SetupAdminBehaviorDB setup admin behavior log db
func SetupAdminBehaviorDB(sqlHandle string) error {
	mysql.Instance().DB(sqlHandle).DropTableIfExists(&models.AdminBehavior{})
	err := mysql.Instance().DB(sqlHandle).
		Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&models.AdminBehavior{}).Error
	if err != nil {
		return err
	}
	return nil
}
