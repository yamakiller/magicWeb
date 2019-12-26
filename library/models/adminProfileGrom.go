package models

import "time"

//AdminProfile perm group info
type AdminProfile struct {
	ID        string `grom:"primary_key;type:varchar(36)"`
	Name      string `grom:"type:varchar(16);not null;"`
	Data      string `grom:"type:longtext"`
	Order     uint8  `grom:"default 0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
