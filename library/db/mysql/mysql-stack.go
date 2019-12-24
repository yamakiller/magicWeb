package mysql

import (
	"sync"

	"github.com/jinzhu/gorm"

	"github.com/yamakiller/magicLibs/dbs"
)

var (
	oneSQLStack sync.Once
	stack       *SQLStack
)

//Instance doc
// @Summary mysql stack instance
// @Method Instance
// @Return (*MySQLStack)
func Instance() *SQLStack {
	oneSQLStack.Do(func() {
		stack = &SQLStack{_cs: make(map[string]*SQLDDB)}
	})
	return stack
}

// SQLDDB doc
// @Summary mysql handle
// @Struct MySQLms
// @Inherit dbs.MySQLDB
// @Member  bool  is close ?
type SQLDDB struct {
	dbs.MySQLGORM
	_closed bool
}

// SQLStack doc
// @Summary mysql client
// @Struct MySQLStack doc
// @Member (*dbs.MySQLDB)
type SQLStack struct {
	_cs map[string]*SQLDDB
}

//DB Return grom db object
func (slf *SQLStack) DB(key string) *gorm.DB {
	if v, ok := slf._cs[key]; ok {
		return v.DB()
	}
	return nil
}

// Append doc
// @Summary config to Append mysql pool handle
// @Param (*dbs.MySQLDeploy) mysql config
// @Return (*dbs.MySQLDB) create mysql connection pools
// @Return (error)
func (slf *SQLStack) Append(key string, d *dbs.MySQLGormDeploy) (*SQLDDB, error) {
	c := &SQLDDB{}
	e := dbs.DoMySQLGormDeploy(&c.MySQLGORM, d)
	if e != nil {
		return nil, e
	}
	c._closed = false
	slf._cs[key] = c
	return c, nil
}

// AppendObject doc
// @Summary Append mysql pool handle
// @Method Append doc
// @Param (*dbs.MySQLDeploy) mysql config
func (slf *SQLStack) AppendObject(key string, c *SQLDDB) {
	slf._cs[key] = c
}

// IsConnected doc
// @Summary is mysql connected.
// @Method IsConnected doc
// @Return (bool)
func (slf *SQLStack) IsConnected(key string) bool {
	if _, ok := slf._cs[key]; !ok {
		return false
	}
	return true
}

// Close doc
// @Summary close mysql
// @Method Close doc
func (slf *SQLStack) Close() {
	for k, v := range slf._cs {
		if !v._closed {
			v.Close()
			v._closed = true
		}
		delete(slf._cs, k)
	}
}

/*
// Query doc
// @Summary Query data
// @Method Query
// @Param  (string) sql handle key
// @Param  (string) sql
// @Param  (..interface{}) sql args
// @Return (*dbs.MySQLReader)
func (slf *SQLStack) Query(key string, ssql string, args ...interface{}) (*dbs.MySQLReader, error) {
	if v, ok := slf._cs[key]; ok {
		return v.Query(ssql, args...)
	}

	return nil, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

// QueryPage doc
// @Summary Query page data
// @Method  Query
// @Param   (string) sql handle key
// @Param   (string) table files (xxx,xxx)
// @Param   (string) table names (xxx,xxx)
// @Param   (string) query condition
// @Param   (string) query order mode
// @Param   (int) page
// @Param   (int) pageSize
// @Return  (int) pageCount
// @Param   (...interface{}) where args
// @Return  (*dbs.MySQLReader) reader
// @Return  (error) ree
func (slf *SQLStack) QueryPage(key, fileds, tables, where, order string, page, pageSize int, args ...interface{}) (pageCount int, reader *dbs.MySQLReader, err error) {
	if v, ok := slf._cs[key]; ok {
		return v.QueryPage(fileds, tables, where, order, page, pageSize, args...)
	}

	return 0, nil, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

//Insert doc
// @Summary insert data
// @Method Insert
// @Param  (string) sql handle key
// @Param  (string) sql
// @Param  (..interface{}) sql args
// @Return (int) insert data of number
func (slf *SQLStack) Insert(key string, ssql string, args ...interface{}) (int, error) {
	if v, ok := slf._cs[key]; ok {
		n, e := v.Insert(ssql, args...)
		if e != nil {
			return 0, e
		}

		return int(n), nil
	}
	return 0, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

// Update doc
// @Summary insert data
// @Method Update doc
// @Param  (string) sql handle key
// @Param  (string) sql
// @Param  (..interface{}) sql args
// @Return (int) insert data of number
func (slf *SQLStack) Update(key string, ssql string, args ...interface{}) (int, error) {
	if v, ok := slf._cs[key]; ok {
		n, e := v.Update(ssql, args...)
		if e != nil {
			return 0, e
		}

		return int(n), nil
	}

	return 0, fmt.Errorf("Non-existent %s MySQL Connect", key)
}

// Close doc
// @Summary close mysql
// @Method Close doc
func (slf *SQLStack) Close() {
	for k, v := range slf._cs {
		if !v._closed {
			v.Close()
			v._closed = true
		}
		delete(slf._cs, k)
	}
}
*/
