package models

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusInvited     UserStatus = "INVITED"
	UserStatusIncomplete  UserStatus = "INCOMPLETE"
	UserStatusComplete    UserStatus = "COMPLETE"
)

type UserRole string

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email          string    `gorm:"unique;not null" json:"email"`
	ProfilePicture *string   `json:"profilePicture"`
	CreatedAt      time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"not null" json:"updatedAt"`
	LastLogin      *time.Time `json:"lastLogin,omitempty"`
	Status         UserStatus `gorm:"not null" json:"status"`
	// Relationships
	// Activity      []Activity      `gorm:"foreignKey:CreatedBy" json:"activity,omitempty"`
	// Teams         []UserTeam      `gorm:"foreignKey:UserId" json:"teams,omitempty"`
	Workspaces    []UserWorkspace `gorm:"foreignKey:UserId" json:"workspaces,omitempty"`
	AssignedTasks []Task          `gorm:"foreignKey:AssignedTo" json:"assignedTasks,omitempty"`
}

type UserWorkspace struct {
	UserId      uuid.UUID `gorm:"primaryKey" json:"userId"`
	WorkspaceId uuid.UUID `gorm:"primaryKey" json:"workspaceId"`
	UserName    string    `gorm:"not null" json:"userName"`
	Role        UserRole  `gorm:"not null" json:"role"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
	// Relationships
	User      User      `gorm:"foreignKey:UserId" json:"user"`
	// Workspace Workspace `gorm:"foreignKey:WorkspaceId" json:"workspace"`
}

type CreateUserPayload struct {
	Email	string `json:"email"`
}

type UpdateUserPayload struct {
	ProfilePicture	string `json:"profilePicture"`
}
