package database

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/yamakiller/magicLibs/dbs"
	"github.com/yamakiller/magicWeb/library/db/mysql"
)

//SQLDeploys mysql deploy group
type SQLDeploys struct {
	Items []SQLDeployItem `xml:"items" yaml:"items" json:"items"`
}

//SQLDeployItem mysql deploy item
type SQLDeployItem struct {
	Key string `xml:"key" yaml:"key" json:"key"`
	dbs.MySQLGormDeploy
}

//SQLRegister register mysql connection pools
func SQLRegister(project string, config *SQLDeploys) error {
	for _, v := range config.Items {
		keys := strings.Split(v.Key, "|")
		conn, err := mysql.Instance().Append(keys[0], &v.MySQLGormDeploy)
		if err != nil {
			return err
		}

		for i := 1; i < len(keys); i++ {
			mysql.Instance().AppendObject(keys[1], conn)
		}
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return project + "_" + defaultTableName
	}
	return nil
}
