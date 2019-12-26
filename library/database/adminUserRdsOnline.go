package database

import (
	"time"

	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/db/redis"
)

const (
	//OnlineUserTokenKey  user online redis hash token field name
	onlineUserTokenKey = "token"
	//OnlineUserAccountKey  user online redis hash user name field name
	onlineUserAccountKey = "account"
	//OnlineUserPwdKey  user online redis hash user password field name
	onlineUserPwdKey = "pwd"
	//OnlineUserSecretKey  user online redis hash user secret field name
	onlineUserSecretKey = "secret"
	//OnlineUserProfileKey  user online redis hash user permissions field name
	onlineUserProfileKey = "profile"
	//OnlineUserFeatureKey  user online redis hash user feature field name [0.nomal user 1.admin user]
	onlineUserBackstageKey = "backstage"
	//OnlineUserLasttime    user online redis hash user logined last time field name
	onlineUserLasttimeKey = "lasttime"
	//onlineUserActived     user online redis hash user actived last time filed name
	onlineUserActivedKey = "actived"
)

//CreateRdsOnlineAdminUserVal doc
//Summary Setting Online User data
//Method SetOnlineUser
//Param (string) user id
//Param (string) user name
//Param (string) user secret
//Param (string) user perm
//Param (string) user last login time
//Param (int)    user backstage 0.nomal user 1.admin user
//Param (int)    user online expire time /mintue
//Return (error)
func CreateRdsOnlineAdminUserVal(db int,
	userid, token, username, userpwd, secret, profile, lasttime string,
	backstage, expire int) error {
	if _, err := redis.Instance().Do(db, "HMSET", common.GetRdsOnlineKey(userid),
		onlineUserTokenKey, token,
		onlineUserAccountKey, username,
		onlineUserPwdKey, userpwd,
		onlineUserSecretKey, secret,
		onlineUserProfileKey, profile,
		onlineUserBackstageKey, backstage,
		onlineUserLasttimeKey, lasttime,
		onlineUserActivedKey, time.Now().UnixNano()/int64(time.Millisecond)); err != nil {
		redis.Instance().Do(db, "DEL", common.GetRdsOnlineKey(userid))
		return err
	}

	if _, err := redis.Instance().Do(db, "expire", common.GetRdsOnlineKey(userid), expire); err != nil {
		RemoveOnlineAdminUser(db, userid)
		return err
	}

	return nil
}

//WithRdsOnlineAdminToken Update Online User Token
//Param  (int) db
//Param  (string) user id
//Param  (string) user token
//Return (error)
func WithRdsOnlineAdminToken(db int, userid, token string) error {
	return withRdsOnlineVal(db, userid, onlineUserTokenKey, token)
}

//WithRdsOnlineAdminActived Update Online User Actived last time
//Param  (int) db
//Param  (string) user id
//Return (error)
func WithRdsOnlineAdminActived(db int, userid string) error {
	return withRdsOnlineVal(db, userid, onlineUserActivedKey, time.Now().UnixNano()/int64(time.Millisecond))
}

//WithRdsOnlineAdminProfile Update Online User profile informat
//Param  (int) db
//Param  (string) user id
//Param  (string) user profile
//Return (error)
func WithRdsOnlineAdminProfile(db int, userid, profile string) error {
	return withRdsOnlineVal(db, userid, onlineUserProfileKey, profile)
}

//WithRdsOnlineAdminBackstage Update Online User backstate state
//Param  (int) db
//Param  (string) user id
//Param  (string) user backstate
//Return (error)
func WithRdsOnlineAdminBackstage(db int, userid string, backstate int) error {
	return withRdsOnlineVal(db, userid, onlineUserBackstageKey, backstate)
}

//VerifyRdsOnlineAdminAccount verfiy Online User Account is exists
//Param  (int) db
//Param  (string) user id
//Return (bool)
//Return (error)
func VerifyRdsOnlineAdminAccount(db int, userid string) (bool, error) {
	return verifyRdsOnlineUserVal(db, userid, onlineUserAccountKey)
}

//GetRdsOnlineAdminAccount return Online user Account
func GetRdsOnlineAdminAccount(db int, userid string) (string, error) {
	act, err := getRdsOnlineUserVal(db, userid, onlineUserAccountKey)
	if err != nil {
		return "", err
	}

	return act.(string), nil
}

//GetRdsOnlineAdminPassword return Online user Password
func GetRdsOnlineAdminPassword(db int, userid string) (string, error) {
	pwd, err := getRdsOnlineUserVal(db, userid, onlineUserPwdKey)
	if err != nil {
		return "", err
	}

	return pwd.(string), nil
}

//GetRdsOnlineAdminSecret return Online user secret
func GetRdsOnlineAdminSecret(db int, userid string) (string, error) {
	secret, err := getRdsOnlineUserVal(db, userid, onlineUserSecretKey)
	if err != nil {
		return "", err
	}

	return secret.(string), nil
}

//GetRdsOnlineAdminProfile return Online user profile
func GetRdsOnlineAdminProfile(db int, userid string) (string, error) {
	profile, err := getRdsOnlineUserVal(db, userid, onlineUserProfileKey)
	if err != nil {
		return "", err
	}

	return profile.(string), nil
}

//GetRdsOnlineAdminBackstage return Online user backstate state
func GetRdsOnlineAdminBackstage(db int, userid string) (int, error) {
	backstate, err := getRdsOnlineUserVal(db, userid, onlineUserBackstageKey)
	if err != nil {
		return 0, err
	}

	return backstate.(int), nil
}

//GetRdsOnlineAdminLoginLastTime return Online user logined last time
func GetRdsOnlineAdminLoginLastTime(db int, userid string) (string, error) {
	lasttime, err := getRdsOnlineUserVal(db, userid, onlineUserLasttimeKey)
	if err != nil {
		return "", err
	}

	return lasttime.(string), nil
}

//GetRdsOnlineAdminActived return Online user actived last time
func GetRdsOnlineAdminActived(db int, userid string) (int64, error) {
	lasttime, err := getRdsOnlineUserVal(db, userid, onlineUserActivedKey)
	if err != nil {
		return 0, err
	}

	return lasttime.(int64), nil
}

func withRdsOnlineVal(db int, userid string, key string, value interface{}) error {
	if _, err := redis.Instance().Do(db, "HMSet", common.GetRdsOnlineKey(userid),
		key, value); err != nil {
		return err
	}
	return nil
}

func verifyRdsOnlineUserVal(db int, userid, key string) (bool, error) {
	v, err := redis.Instance().Do(db, "HEXISTS", common.GetRdsOnlineKey(userid), key)
	if err != nil {
		return false, err
	}

	return v.(bool), nil
}

//GetRdsOnlineUserVal doc
//Summary Query Online User Key => Value
//Param (string) user id
//Param (string) key
//Return (interface{}) value
//Return (error)
func getRdsOnlineUserVal(db int, userid, key string) (interface{}, error) {
	return redis.Instance().Do(db, "HGET", common.GetRdsOnlineKey(userid), key)
}

//WithRdsOnlineAdminExpire doc
//Summary Update Online user data expire
//Param (string) user id
//Param  (int) expire second
func WithRdsOnlineAdminExpire(db int, userid string, expire int) error {
	if _, err := redis.Instance().Do(db, "expire", common.GetRdsOnlineKey(userid), expire); err != nil {
		return err
	}

	return nil
}

//RemoveOnlineAdminUser doc
//Summary Remove Online User data
//Method RemoveOnlineUser
//Param (string) user id
//Param (string) user name
//Param (string) user secret
//Param (string) user perm
//Param (string) user last login time
//Param (int)    user feature 0.nomal user 1.admin user
//Return (error)
func RemoveOnlineAdminUser(db int, userid string) error {
	if _, err := redis.Instance().Do(db,
		"HDEL", common.GetRdsOnlineKey(userid),
		onlineUserTokenKey,
		onlineUserAccountKey,
		onlineUserPwdKey,
		onlineUserSecretKey,
		onlineUserProfileKey,
		onlineUserBackstageKey,
		onlineUserLasttimeKey,
		onlineUserActivedKey); err != nil {
		return err
	}

	return nil
}
