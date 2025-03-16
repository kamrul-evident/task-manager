package models

// UserRole enum
type UserRole string

const (
	Admin UserRole = "admin"
	Empoyee UserRole = "employee"
	Manager UserRole = "manager"
	Other UserRole = "other"
	SoftwareEngineer UserRole = "software_engineer"
)

// User model
type User struct {
	BaseModelWithUUID
	Email string `gorm:"uniqueIndex;size:128;not null" json:"email"`
	Password string `gorm:"not null" json:"-"` // Exclude from JSON
	IsActive bool `gorm:"default:true" json:"is_active"`
	IsAdmin bool `'gorm:"default:false" json:"is_admin"`
	Role UserRole `gorm:"default:other" json:"role"`
	Tasks []Task `gorm:"foreignKey:AssigneeID" json:"tasks"`
}

