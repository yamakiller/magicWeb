package database

import (
	"time"

	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminProfileQueryAll Retusn admin profile group all
func AdminProfileQueryAll(sqlHandle string) (profiles []models.AdminProfile, err error) {
	err = mysql.Instance().
		DB(sqlHandle).
		Where("deleted_at IS NOT NULL").
		Order("CreatedAt asc").Find(profiles).Error
	if err != nil {
		return
	}

	return
}

//AdminProfileQuery Return page admin profile group and total
func AdminProfileQuery(sqlHandle string, page, pageSize int) (profiles []models.AdminProfile, total int, err error) {
	err = mysql.Instance().
		DB(sqlHandle).
		Where("deleted_at IS NOT NULL").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("CreatedAt asc").Find(profiles).Error
	if err != nil {
		return
	}

	mysql.Instance().DB(sqlHandle).Model(&models.AdminProfile{}).Count(&total)
	return
}

//AdminProfileAppend append a profile group data
func AdminProfileAppend(sqlHandle string, profile *models.AdminProfile) (string, error) {
	profile.ID = util.SpawnUUID()
	return profile.ID, mysql.Instance().DB(sqlHandle).Create(profile).Error
}

//AdminProfileUpdate update a profile group data
func AdminProfileUpdate(sqlHandle string, profile *models.AdminProfile) error {
	return mysql.Instance().DB(sqlHandle).Update(profile).Error
}

//AdminProfileDelete delete a profile  group data
func AdminProfileDelete(sqlHandle, id string) error {
	return mysql.Instance().DB(sqlHandle).Where("id=?", id).Update("deleted_at", time.Now()).Error
}

//AdminProfileQueryName Returns profile group name
func AdminProfileQueryName(sqlHandle string, id string) (string, error) {
	profile := models.AdminProfile{}
	if err := mysql.Instance().DB(sqlHandle).Select("name").Where("id = ?", id).First(&profile).Error; err != nil {
		return "", err
	}

	return profile.Name, nil
}
