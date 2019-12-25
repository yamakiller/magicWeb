package models

import (
	"github.com/jinzhu/gorm"
)

//Profile perm group info
type Profile struct {
	gorm.Model
	ID    string `grom:"primary_key;type:varchar(36)"`
	Name  string `grom:"type:varchar(16);not null;"`
	Data  string `grom:"size:4096"`
	Order uint8  `grom:"default 0"`
}
