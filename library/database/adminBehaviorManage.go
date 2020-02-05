package database

import (
	"time"

	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminBehaviorAppend append admin behavior log
func AdminBehaviorAppend(sqlHandle, adminUserID, bev string) error {
	behavior := models.AdminBehavior{ID: util.SpawnUUID(),
		Operator:  adminUserID,
		Behavior:  bev,
		CreatedAt: time.Now()}
	return mysql.Instance().DB(sqlHandle).Create(&behavior).Error
}

//AdminBehaviorQuery query admin behavior log
func AdminBehaviorQuery(sqlHandle string,
	page, pageSize int,
	Where string,
	args ...interface{}) (behaviors []models.AdminBehavior, total int, err error) {

	err = mysql.Instance().DB(sqlHandle).
		Where(Where, args...).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("CreatedAt desc").Find(&behaviors).Error
	if err != nil {
		return
	}

	mysql.Instance().DB(sqlHandle).Model(&models.AdminBehavior{}).Count(&total)
	return
}
