package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
	UUID      uuid.UUID  `gorm:"type:uuid;" json:"uuid"`
	OwnerID   string     `json:"owner_id"`
	CreatedAt time.Time  `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time  `orm:"auto_now;type(datetime)" json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("UUID", uuid)
}
