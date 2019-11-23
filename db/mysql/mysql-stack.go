package db

import (
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
		stack = &MySQLStack{}
	})
	return stack
}

//MySQLStack desc
//@struct MySQLStack desc: mysql client
//@member (*dbs.MySQLDB)
type MySQLStack struct {
	_c *dbs.MySQLDB
}

//Initial desc
//@method Initial desc: Initialization mysql pools
//@param (*dbs.MySQLDeploy) mysql config
func (slf *MySQLStack) Initial(d *dbs.MySQLDeploy) error {
	slf._c = &dbs.MySQLDB{}
	e := dbs.DoMySQLDeploy(slf._c, d)
	if e != nil {
		slf._c = nil
		return e
	}
	return nil
}

//IsConnected desc
//@method IsConnected desc: is mysql connected.
//@return (bool)
func (slf *MySQLStack) IsConnected() bool {
	if slf._c == nil {
		return false
	}
	return true
}

//Query desc
//@method Query desc: Query sql
//@param  (string) sql
//@param  (..interface{}) sql args
//@return (*dbs.MySQLReader)
func (slf *MySQLStack) Query(ssql string, args ...interface{}) *dbs.MySQLReader {
	return slf.Query(ssql, args...)
}

//Insert desc
//@method Insert desc: insert data
//@param  (string) sql
//@param  (..interface{}) sql args
//@return (int) insert data of number
func (slf *MySQLStack) Insert(ssql string, args ...interface{}) int {
	n, e := slf._c.Insert(ssql, args...)
	if e != nil {
		return 0
	}

	return int(n)
}

//Update desc
//@method Update desc: insert data
//@param  (string) sql
//@param  (..interface{}) sql args
//@return (int) insert data of number
func (slf *MySQLStack) Update(ssql string, args ...interface{}) int {
	n, e := slf._c.Update(ssql, args...)
	if e != nil {
		return 0
	}
	return int(n)
}

//Close desc
//@method Close desc: close mysql
func (slf *MySQLStack) Close() {
	if slf._c == nil {
		return
	}
	slf._c.Close()
	slf._c = nil
}
