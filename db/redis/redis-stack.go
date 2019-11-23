package db

import (
	"sync"

	"github.com/yamakiller/magicLibs/dbs"
)

var (
	oneRedis sync.Once
	stack    *RedisStack
)

//Instance desc
//@method Instance desc: Reids instance
//@return (*RedisStack)
func Instance() *RedisStack {
	oneRedis.Do(func() {
		stack = &RedisStack{}
	})
	return stack
}

//RedisStack desc
//@struct RedisStack desc: redis client
type RedisStack struct {
	_c *dbs.RedisDB
}

//Initial desc
//@method Initial desc: Initialization redis pools
//@param (*dbs.RedisDeppoy) redis config
func (slf *RedisStack) Initial(d *dbs.RedisDeploy) error {
	slf._c = &dbs.RedisDB{}
	e := dbs.DoRedisDeploy(slf._c, d)
	if e != nil {
		slf._c = nil
		return e
	}
	return nil
}

//IsConnected desc
//@method IsConnected desc: redis is connected
func (slf *RedisStack) IsConnected() bool {
	if slf._c == nil {
		return false
	}
	return true
}

//Do desc
//@method Do desc: execute redis command
//@param (string) command name
//@param (...interface{}) command params
//@return (interface{}) execute result
//@return (error) if execute fail return error, execute success return nil
func (slf *RedisStack) Do(commandName string, args ...interface{}) (interface{}, error) {
	return slf._c.Do(commandName, args...)
}

//Close desc
//@method Close desc: close redis db operation
func (slf *RedisStack) Close() {
	slf.Close()
}
