package models

import "time"

//AdminBehavior admin user behavior
type AdminBehavior struct {
	ID         string    `gorm:"primary_key;type:varchar(36);not null;"`
	OperUserID string    `gorm:"type:varchar(36);index;not null;"`
	Behavior   string    `gorm:"type:varchar(1024);not null;"`
	CreatedAt  time.Time `gorm:"index:idx_behavior_time;not null;"`
}
