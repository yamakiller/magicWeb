package models

import "time"

//AdminUser Table
type AdminUser struct {
	ID              string       `gorm:"primary_key;type:varchar(36);not null;"`
	Account         string       `gorm:"type:varchar(32);not null;index:account_idx;"`
	Password        string       `gorm:"type:varchar(64);not null;"`
	Secret          string       `gorm:"type:varchar(16);not null;"`
	Nick            string       `gorm:"type:varchar(16);not null;"`
	Level           int          `gorm:"not null;"`
	Email           string       `gorm:"type:varchar(128);not null;"`
	Mobile          string       `gorm:"type:varchar(20);not null;"`
	Profile         AdminProfile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:ID;"`
	ProfileID       string       `gorm:"type:varchar(36);index;not null;"`
	Backstage       uint8        `gorm:"not null;"`
	State           uint8        `gorm:"not null;"`
	Source          string       `gorm:"type:varchar(64);not null;"`
	Fail            uint16       `gorm:"not null;"`
	Logined         uint16       `gorm:"not null;"`
	LoginedIP       string       `gorm:"type:varchar(32);not null;"`
	CreateIP        string       `gorm:"type:varchar(32);not null;"`
	FailLastTime    time.Time    `gorm:"not null;"`
	LoginedLastTime time.Time    `gorm:"not null;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
