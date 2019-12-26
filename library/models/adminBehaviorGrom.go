package models

import "time"

//AdminBehavior admin user behavior
type AdminBehavior struct {
	ID          string    `grom:"primary_key;type:varchar(36)"`
	OperAdminID string    `grom:"type:varchar(36);index"`
	Behavior    string    `grom:"type:varchar(1024);not null;"`
	CreatedAt   time.Time `grom:"index:idx_behavior_time"`
}
