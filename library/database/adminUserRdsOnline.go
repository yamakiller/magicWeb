package database

import (
	"errors"
	"time"

	"github.com/yamakiller/magicWeb/library/common"
	"github.com/yamakiller/magicWeb/library/db/redis"
)

var (
	//ErrOnlineUserEmpty user not online
	ErrOnlineUserEmpty = errors.New("Online user empty")
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
	//onlineUserRoleKey  user online redis hash user permissions field name
	onlineUserRoleKey = "role"
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
	userid,
	token,
	username,
	userpwd,
	secret,
	role,
	lasttime string,
	backstage,
	expireSec int) error {
	if _, err := redis.Instance().Do(db, "HMSET", common.GetAdminRdsOnlineKey(userid),
		onlineUserTokenKey, token,
		onlineUserAccountKey, username,
		onlineUserPwdKey, userpwd,
		onlineUserSecretKey, secret,
		onlineUserRoleKey, role,
		onlineUserBackstageKey, backstage,
		onlineUserLasttimeKey, lasttime,
		onlineUserActivedKey, time.Now().UnixNano()/int64(time.Millisecond)); err != nil {
		redis.Instance().Do(db, "DEL", common.GetAdminRdsOnlineKey(userid))
		return err
	}

	if _, err := redis.Instance().Do(db, "expire", common.GetAdminRdsOnlineKey(userid),
		expireSec); err != nil {
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
func WithRdsOnlineAdminToken(db int, userid, token string, expireSec int) error {
	if err := withRdsOnlineVal(db, userid, onlineUserTokenKey, token); err != nil {
		return err
	}

	if _, err := redis.Instance().Do(db, "expire", common.GetAdminRdsOnlineKey(userid), expireSec); err != nil {
		//RemoveOnlineAdminUser(db, userid)
		return err
	}
	return nil
}

//WithRdsOnlineAdminActived Update Online User Actived last time
//Param  (int) db
//Param  (string) user id
//Return (error)
func WithRdsOnlineAdminActived(db int, userid string) error {
	return withRdsOnlineVal(db, userid, onlineUserActivedKey, time.Now().UnixNano()/int64(time.Millisecond))
}

//WithRdsOnlineAdminRole Update Online User role informat
//Param  (int) db
//Param  (string) user id
//Param  (string) user role
//Return (error)
func WithRdsOnlineAdminRole(db int, userid, role string) error {
	return withRdsOnlineVal(db, userid, onlineUserRoleKey, role)
}

//WithRdsOnlineAdminBackstage Update Online User backstate state
//Param  (int) db
//Param  (string) user id
//Param  (string) user backstate
//Return (error)
func WithRdsOnlineAdminBackstage(db int, userid string, backstate int) error {
	return withRdsOnlineVal(db, userid, onlineUserBackstageKey, backstate)
}

//WithRdsOnlineAdminPwd Update Online User password
func WithRdsOnlineAdminPwd(db int, userid string, pwd string) error {
	return withRdsOnlineVal(db, userid, onlineUserPwdKey, pwd)
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
	} else if act == nil {
		return "", ErrOnlineUserEmpty
	}

	return act.(string), nil
}

//GetRdsOnlineAdminPassword return Online user Password
func GetRdsOnlineAdminPassword(db int, userid string) (string, error) {
	pwd, err := getRdsOnlineUserVal(db, userid, onlineUserPwdKey)
	if err != nil {
		return "", err
	} else if pwd == nil {
		return "", ErrOnlineUserEmpty
	}

	return pwd.(string), nil
}

//GetRdsOnlineAdminSecret return Online user secret
func GetRdsOnlineAdminSecret(db int, userid string) (string, error) {
	secret, err := getRdsOnlineUserVal(db, userid, onlineUserSecretKey)
	if err != nil {
		return "", err
	} else if secret == nil {
		return "", ErrOnlineUserEmpty
	}

	return secret.(string), nil
}

//GetRdsOnlineAdminRole return Online user role
func GetRdsOnlineAdminRole(db int, userid string) (string, error) {
	role, err := getRdsOnlineUserVal(db, userid, onlineUserRoleKey)
	if err != nil {
		return "", err
	} else if role == nil {
		return "", ErrOnlineUserEmpty
	}

	return role.(string), nil
}

//GetRdsOnlineAdminBackstage return Online user backstate state
func GetRdsOnlineAdminBackstage(db int, userid string) (int, error) {
	backstate, err := getRdsOnlineUserVal(db, userid, onlineUserBackstageKey)
	if err != nil {
		return 0, err
	} else if backstate == nil {
		return 0, ErrOnlineUserEmpty
	}

	return backstate.(int), nil
}

//GetRdsOnlineAdminLoginLastTime return Online user logined last time
func GetRdsOnlineAdminLoginLastTime(db int, userid string) (string, error) {
	lasttime, err := getRdsOnlineUserVal(db, userid, onlineUserLasttimeKey)
	if err != nil {
		return "", err
	} else if lasttime == nil {
		return "", ErrOnlineUserEmpty
	}

	return lasttime.(string), nil
}

//GetRdsOnlineAdminActived return Online user actived last time
func GetRdsOnlineAdminActived(db int, userid string) (int64, error) {
	lasttime, err := getRdsOnlineUserVal(db, userid, onlineUserActivedKey)
	if err != nil {
		return 0, err
	} else if lasttime == nil {
		return 0, ErrOnlineUserEmpty
	}

	return lasttime.(int64), nil
}

func withRdsOnlineVal(db int, userid string, key string, value interface{}) error {
	if _, err := redis.Instance().Do(db, "HMSet", common.GetAdminRdsOnlineKey(userid),
		key, value); err != nil {
		return err
	}
	return nil
}

func verifyRdsOnlineUserVal(db int, userid, key string) (bool, error) {
	v, err := redis.Instance().Do(db, "HEXISTS", common.GetAdminRdsOnlineKey(userid), key)
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
	return redis.Instance().Do(db, "HGET", common.GetAdminRdsOnlineKey(userid), key)
}

//WithRdsOnlineAdminExpire doc
//Summary Update Online user data expire
//Param (string) user id
//Param  (int) expire second
func WithRdsOnlineAdminExpire(db int, userid string, expireSec int) error {
	if _, err := redis.Instance().Do(db, "expire",
		common.GetAdminRdsOnlineKey(userid), expireSec); err != nil {
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
		"HDEL", common.GetAdminRdsOnlineKey(userid),
		onlineUserTokenKey,
		onlineUserAccountKey,
		onlineUserPwdKey,
		onlineUserSecretKey,
		onlineUserRoleKey,
		onlineUserBackstageKey,
		onlineUserLasttimeKey,
		onlineUserActivedKey); err != nil {
		return err
	}

	return nil
}
