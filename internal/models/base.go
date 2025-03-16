package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base Model with UUID provides common fields
type BaseModelWithUUID struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UID       string         `gorm:"uniqueIndex;size:36" json:"uid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"` // Nullable, no default on create
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Before Create hook to set the UID and timestamps
func (b *BaseModelWithUUID) BeforeCreate(tx *gorm.DB) error {
	if b.UID == "" {
		b.UID = uuid.New().String()
	}
	b.CreatedAt = time.Now()
	return nil
}

// Before Update hook to set UpdatedAt
func (b *BaseModelWithUUID) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now()
	b.UpdatedAt = &now
	return nil
}
