package database

import (
	"github.com/yamakiller/magicLibs/dbs"
	"github.com/yamakiller/magicWeb/library/db/redis"
)

//RedisDeployGroup redis connection pool config group
type RedisDeployGroup struct {
	Items []dbs.RedisDeploy
}

//RedisRegister register redis connection pool
func RedisRegister(config *RedisDeployGroup) error {
	for _, v := range config.Items {
		if err := redis.Instance().Append(&v); err != nil {
			redis.Instance().Close()
			return err
		}
	}
	return nil
}
