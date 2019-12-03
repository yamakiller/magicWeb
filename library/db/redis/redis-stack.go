package redis

import (
	"fmt"
	"sync"

	"github.com/yamakiller/magicLibs/dbs"
)

var (
	oneRedis sync.Once
	stack    *RedisStack
)

//Instance desc
//@Method Instance desc: Reids instance
//@Return (*RedisStack)
func Instance() *RedisStack {
	oneRedis.Do(func() {
		stack = &RedisStack{make(map[int]*dbs.RedisDB)}
	})
	return stack
}

//RedisStack desc
//@Struct RedisStack desc: redis client
//@Member (*dbs.RedisDB)
type RedisStack struct {
	_cs map[int]*dbs.RedisDB
}

//Append desc
//@Method Append desc: Append redis pools
//@Param (*dbs.RedisDeppoy) redis config
func (slf *RedisStack) Append(d *dbs.RedisDeploy) error {
	if _, ok := slf._cs[d.DB]; ok {
		return fmt.Errorf("redis %d already exists", d.DB)
	}

	c := &dbs.RedisDB{}
	e := dbs.DoRedisDeploy(c, d)
	if e != nil {
		return e
	}

	slf._cs[d.DB] = c
	return nil
}

//IsConnected desc
//@Method IsConnected desc: redis is connected
//@Param (int) db
func (slf *RedisStack) IsConnected(db int) bool {
	if _, ok := slf._cs[db]; !ok {
		return false
	}
	return true
}

//Do desc
//@Method Do desc: execute redis command
//@Param (int)    db
//@Param (string) command name
//@Param (...interface{}) command params
//@Return (interface{}) execute result
//@Return (error) if execute fail return error, execute success return nil
func (slf *RedisStack) Do(db int, commandName string, args ...interface{}) (interface{}, error) {
	if _, ok := slf._cs[db]; !ok {
		return nil, fmt.Errorf("redis %d already exists", db)
	}

	c := slf._cs[db]

	return c.Do(commandName, args...)
}

//Close desc
//@Method Close desc: close redis db operation
func (slf *RedisStack) Close() {
	for k, v := range slf._cs {
		v.Close()
		delete(slf._cs, k)
	}
}
