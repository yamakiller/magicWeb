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

//CreateRdsOnlineUserVal doc
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
func CreateRdsOnlineUserVal(db int,
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
		RemoveOnlineUser(db, userid)
		return err
	}

	return nil
}

//WithRdsOnlineToken Update Online User Token
//Param  (int) db
//Param  (string) user id
//Param  (string) user token
//Return (error)
func WithRdsOnlineToken(db int, userid, token string) error {
	return withRdsOnlineVal(db, userid, onlineUserTokenKey, token)
}

//WithRdsOnlineActived Update Online User Actived last time
//Param  (int) db
//Param  (string) user id
//Return (error)
func WithRdsOnlineActived(db int, userid string) error {
	return withRdsOnlineVal(db, userid, onlineUserActivedKey, time.Now().UnixNano()/int64(time.Millisecond))
}

//WithRdsOnlineProfile Update Online User profile informat
//Param  (int) db
//Param  (string) user id
//Param  (string) user profile
//Return (error)
func WithRdsOnlineProfile(db int, userid, profile string) error {
	return withRdsOnlineVal(db, userid, onlineUserProfileKey, profile)
}

//WithRdsOnlineBackstage Update Online User backstate state
//Param  (int) db
//Param  (string) user id
//Param  (string) user backstate
//Return (error)
func WithRdsOnlineBackstage(db int, userid string, backstate int) error {
	return withRdsOnlineVal(db, userid, onlineUserBackstageKey, backstate)
}

//VerifyRdsOnlineAccount verfiy Online User Account is exists
//Param  (int) db
//Param  (string) user id
//Return (bool)
//Return (error)
func VerifyRdsOnlineAccount(db int, userid string) (bool, error) {
	return verifyRdsOnlineUserVal(db, userid, onlineUserAccountKey)
}

//GetRdsOnlineAccount return Online user Account
func GetRdsOnlineAccount(db int, userid string) (string, error) {
	act, err := getRdsOnlineUserVal(db, userid, onlineUserAccountKey)
	if err != nil {
		return "", err
	}

	return act.(string), nil
}

//GetRdsOnlinePassword return Online user Password
func GetRdsOnlinePassword(db int, userid string) (string, error) {
	pwd, err := getRdsOnlineUserVal(db, userid, onlineUserPwdKey)
	if err != nil {
		return "", err
	}

	return pwd.(string), nil
}

//GetRdsOnlineSecret return Online user secret
func GetRdsOnlineSecret(db int, userid string) (string, error) {
	secret, err := getRdsOnlineUserVal(db, userid, onlineUserSecretKey)
	if err != nil {
		return "", err
	}

	return secret.(string), nil
}

//GetRdsOnlineProfile return Online user profile
func GetRdsOnlineProfile(db int, userid string) (string, error) {
	profile, err := getRdsOnlineUserVal(db, userid, onlineUserProfileKey)
	if err != nil {
		return "", err
	}

	return profile.(string), nil
}

//GetRdsOnlineBackstage return Online user backstate state
func GetRdsOnlineBackstage(db int, userid string) (int, error) {
	backstate, err := getRdsOnlineUserVal(db, userid, onlineUserBackstageKey)
	if err != nil {
		return 0, err
	}

	return backstate.(int), nil
}

//GetRdsOnlineLoginLastTime return Online user logined last time
func GetRdsOnlineLoginLastTime(db int, userid string) (string, error) {
	lasttime, err := getRdsOnlineUserVal(db, userid, onlineUserLasttimeKey)
	if err != nil {
		return "", err
	}

	return lasttime.(string), nil
}

//GetRdsOnlineActived return Online user actived last time
func GetRdsOnlineActived(db int, userid string) (int64, error) {
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

//WithRdsOnlineExpire doc
//Summary Update Online user data expire
//Param (string) user id
//Param  (int) expire second
func WithRdsOnlineExpire(db int, userid string, expire int) error {
	if _, err := redis.Instance().Do(db, "expire", common.GetRdsOnlineKey(userid), expire); err != nil {
		return err
	}

	return nil
}

//RemoveOnlineUser doc
//Summary Remove Online User data
//Method RemoveOnlineUser
//Param (string) user id
//Param (string) user name
//Param (string) user secret
//Param (string) user perm
//Param (string) user last login time
//Param (int)    user feature 0.nomal user 1.admin user
//Return (error)
func RemoveOnlineUser(db int, userid string) error {
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
