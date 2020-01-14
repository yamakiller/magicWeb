package handler

import (
	"fmt"
	"time"

	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
	"iov.bapi/app/config"
)

//SetupAdminUserDB setup admin user db
func SetupAdminUserDB(sqlHandle string) error {
	mysql.Instance().DB(sqlHandle).DropTableIfExists(&models.AdminUser{})
	err := mysql.Instance().DB(config.SQLUHandle).Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&models.AdminUser{}).Error
	if err != nil {
		return err
	}

	adminUser := models.AdminUser{}
	adminUser.ID = util.SpawnUUID()

	pwd := "admin"
	secret := util.RandStr(16)
	password, e := util.AesEncrypt(secret, pwd)
	if e != nil {
		return fmt.Errorf("install admin user password encrypt error:%+v", e)
	}

	adminUser.Account = "admin"
	adminUser.Password = password
	adminUser.Nick = "超级管理员"
	adminUser.Secret = secret
	adminUser.State = 0
	adminUser.Source = "local"
	adminUser.Backstage = 1
	adminUser.Roles = "['admin']"
	adminUser.FailLastTime = time.Time{}
	adminUser.LoginedLastTime = time.Now()
	adminUser.LoginedIP = "localhost"
	adminUser.CreatedAt = time.Now()
	adminUser.UpdatedAt = adminUser.CreatedAt

	if err := mysql.Instance().DB(sqlHandle).Create(adminUser).Error; err != nil {
		return err
	}

	return nil
}
