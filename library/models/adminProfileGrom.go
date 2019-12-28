package models

import "time"

//AdminProfile Table
type AdminProfile struct {
	ID        string `gorm:"primary_key;type:varchar(36);not null;"`
	Name      string `gorm:"type:varchar(16);not null;"`
	Data      string `gorm:"type:longtext;"`
	Order     uint8  `gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
