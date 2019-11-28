package db

import (
	"fmt"
	"sync"

	"github.com/yamakiller/magicLibs/dbs"
)

var (
	oneSQLStack sync.Once
	stack       *MySQLStack
)

//Instance desc
//@method Instance desc: mysql stack instance
//@return (*MySQLStack)
func Instance() *MySQLStack {
	oneSQLStack.Do(func() {
		stack = &MySQLStack{_cs: make(map[string]*dbs.MySQLDB)}
	})
	return stack
}

//MySQLStack desc
//@struct MySQLStack desc: mysql client
//@member (*dbs.MySQLDB)
type MySQLStack struct {
	_cs map[string]*dbs.MySQLDB
}

//Append desc
//@method Append desc: Append mysql pool handle
//@param (*dbs.MySQLDeploy) mysql config
func (slf *MySQLStack) Append(key string, d *dbs.MySQLDeploy) error {
	c := &dbs.MySQLDB{}
	e := dbs.DoMySQLDeploy(c, d)
	if e != nil {
		return e
	}
	slf._cs[key] = c
	return nil
}

//IsConnected desc
//@method IsConnected desc: is mysql connected.
//@return (bool)
func (slf *MySQLStack) IsConnected(key string) bool {
	if _, ok := slf._cs[key]; !ok {
		return false
	}
	return true
}

//Query desc
//@method Query desc: Query sql
//@param  (string) sql handle key
//@param  (string) sql
//@param  (..interface{}) sql args
//@return (*dbs.MySQLReader)
func (slf *MySQLStack) Query(key string, ssql string, args ...interface{}) (*dbs.MySQLReader, error) {
	if v, ok := slf._cs[key]; ok {
		return v.Query(ssql, args...)
	}

	return nil, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

//Insert desc
//@method Insert desc: insert data
//@param  (string) sql handle key
//@param  (string) sql
//@param  (..interface{}) sql args
//@return (int) insert data of number
func (slf *MySQLStack) Insert(key string, ssql string, args ...interface{}) (int, error) {
	if v, ok := slf._cs[key]; ok {
		n, e := v.Insert(ssql, args...)
		if e != nil {
			return 0, e
		}

		return int(n), nil
	}
	return 0, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

//Update desc
//@method Update desc: insert data
//@param  (string) sql handle key
//@param  (string) sql
//@param  (..interface{}) sql args
//@return (int) insert data of number
func (slf *MySQLStack) Update(key string, ssql string, args ...interface{}) (int, error) {
	if v, ok := slf._cs[key]; ok {
		n, e := v.Update(ssql, args...)
		if e != nil {
			return 0, e
		}

		return int(n), nil
	}

	return 0, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

//Close desc
//@method Close desc: close mysql
func (slf *MySQLStack) Close() {
	for k, v := range slf._cs {
		v.Close()
		delete(slf._cs, k)
	}
}
