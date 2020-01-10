package handler

import (
	"time"

	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

const (
	//AdminSuperProfileID Super administrator user group id
	adminSuperProfileID = "9ecb9d8b-ee9-4bc8-b6c9-10874e50a435"
)

//SetupAdminProfileDB Mount user rights group DB
func SetupAdminProfileDB(sqlHandle string) error {
	//1.create table
	mysql.Instance().DB(sqlHandle).DropTableIfExists(&models.AdminProfile{})
	err := mysql.Instance().DB(sqlHandle).
		Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&models.AdminProfile{}).Error
	if err != nil {
		return err
	}
	//2.insert base data
	adminProfile := auth.AdminUserProfileItems{}
	adminProfile.Items = append(adminProfile.Items, auth.AdminUserProfile{URI: "ALL"})
	s := util.JSONSerialize(adminProfile)
	profileData := models.AdminProfile{}
	profileData.ID = adminSuperProfileID
	profileData.Name = "超级管理员"
	profileData.Data = s
	profileData.CreatedAt = time.Now()
	profileData.UpdatedAt = time.Now()
	if err := mysql.Instance().DB(sqlHandle).Create(profileData).Error; err != nil {
		return err
	}

	return nil
}
