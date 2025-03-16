package models

import (
	"time"
)

// Priority enum
type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

// TaskStatus enum
type TaskStatus string

const (
	Assigned  TaskStatus = "assigned"
	Pending   TaskStatus = "pending"
	Ongoing   TaskStatus = "ongoing"
	Completed TaskStatus = "completed"
	Archived  TaskStatus = "archived"
)

// Task model
type Task struct {
	NameSlugDescriptionBaseModel // Fixed: Correct spelling from NameSulgDescriptionBaseModel
	DueDate    *time.Time        `json:"due_date,omitempty"`
	Priority   Priority          `gorm:"default:'low'" json:"priority"`
	Category   *string           `json:"category,omitempty"`
	Status     TaskStatus        `gorm:"default:'pending'" json:"status"`
	AssigneeID *uint             `gorm:"index" json:"assignee_id,omitempty"`
	Assignee   User              `gorm:"foreignKey:AssigneeID" json:"assignee"`
}