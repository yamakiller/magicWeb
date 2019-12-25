package handler

import (
	"github.com/yamakiller/magicLibs/util"
	"github.com/yamakiller/magicWeb/library/auth"
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

const (
	//AdminSuperProfileID Super administrator user group id
	AdminSuperProfileID = "9ecb9d8b-ee9-4bc8-b6c9-10874e50a435"
	//NomalProfileID General user group
	NomalProfileID = "02aab766-afa1-4190-a8bf-6e7a64480d1b"
)

//SetupProfileDB Mount user rights group DB
func SetupProfileDB(sqlHandle string) error {
	//1.create table
	mysql.Instance().DB(sqlHandle).DropTableIfExists(&models.Profile{})
	err := mysql.Instance().DB(sqlHandle).
		Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&models.Profile{}).Error
	if err != nil {
		return err
	}
	//2.insert base data
	adminProfile := auth.UserProfileItems{}
	adminProfile.Items = append(adminProfile.Items, auth.UserProfile{URI: "ALL", Auth: auth.ProfileAll})
	s := util.JSONSerialize(adminProfile)
	profileData := models.Profile{}
	profileData.ID = AdminSuperProfileID
	profileData.Name = "超级管理员"
	profileData.Data = s
	if err := mysql.Instance().DB(sqlHandle).Create(profileData).Error; err != nil {
		return err
	}

	///------------------------------------------------------------------------------------------------------------
	nomalUserProfile := auth.UserProfileItems{}
	nomalUserProfile.Items = append(nomalUserProfile.Items, auth.UserProfile{URI: "None", Auth: 0})
	s = util.JSONSerialize(nomalUserProfile)
	profileData.ID = NomalProfileID
	profileData.Name = "普通用户"
	profileData.Data = s
	if err := mysql.Instance().DB(sqlHandle).Create(profileData).Error; err != nil {
		return err
	}

	return nil
}
