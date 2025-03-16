package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type NameSlugDescriptionBaseModel struct {
	BaseModelWithUUID
	Name        string  `gorm:"index;size:128" json:"name"`
	Slug        string  `gorm:"index" json:"slug"`
	Description *string `json:"description,omitempty"`
}

// BeforeCreate hook to generate slug
func (n *NameSlugDescriptionBaseModel) BeforeCreate(tx *gorm.DB) error {
	n.Slug = slug.Make(n.Name)
	return n.BaseModelWithUUID.BeforeCreate(tx)
}

// BeforeUpdate hook to update slug and call base hook
func (n *NameSlugDescriptionBaseModel) BeforeUpdate(tx *gorm.DB) error {
	n.Slug = slug.Make(n.Name)
	return n.BaseModelWithUUID.BeforeUpdate(tx)
}