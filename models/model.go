package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleBasic UserRole = "BASIC"
)

type TaskPriority string

const (
	TaskPriorityNone    TaskPriority = "NONE"
	TaskPriorityLow     TaskPriority = "LOW"
	TaskPriorityMedium  TaskPriority = "MEDIUM"
	TaskPriorityHigh    TaskPriority = "HIGH"
	TaskPriorityUrgent  TaskPriority = "URGENT"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusReview     TaskStatus = "REVIEW"
	TaskStatusDone       TaskStatus = "DONE"
	TaskStatusCancelled  TaskStatus = "CANCELLED"
	TaskStatusDuplicate  TaskStatus = "DUPLICATE"
	TaskStatusBacklog    TaskStatus = "BACKLOG"
)

type Task struct {
	ID           uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title        string        `gorm:"not null" json:"title"`
	Ref          int           `gorm:"not null" json:"ref"`
	Description  *string       `json:"description,omitempty"`
	Status       TaskStatus    `gorm:"not null" json:"status"`
	Priority     *TaskPriority `json:"priority,omitempty"`
	AssignedTo   *uuid.UUID    `json:"assignedTo,omitempty"`
	DueDate      *time.Time    `json:"dueDate,omitempty"`
	CreatedAt    time.Time     `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time     `gorm:"not null" json:"updatedAt"`
	TeamId       uuid.UUID     `gorm:"not null" json:"teamId"`
	ProjectId    *uuid.UUID    `json:"projectId,omitempty"`
	// Relationships
	AssignedUser User       `gorm:"foreignKey:AssignedTo" json:"assignedUser,omitempty"`
	// Team         Team       `gorm:"foreignKey:TeamId" json:"team"`
	// Project      *Project   `gorm:"foreignKey:ProjectId" json:"project,omitempty"`
	// Activity     []Activity `gorm:"foreignKey:TaskId" json:"activity,omitempty"`
	Labels       []TaskLabel `gorm:"many2many:taskLabelsTasks;" json:"labels,omitempty"`
}

type TaskLabel struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Color       string    `gorm:"not null" json:"color"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
	TeamId      uuid.UUID `gorm:"not null" json:"teamId"`
	// Relationships
	// Team  Team   `gorm:"foreignKey:TeamId" json:"team"`
	Tasks []Task `gorm:"many2many:taskLabelsTasks;" json:"tasks,omitempty"`
}

// type ProjectStatus string

// const (
// 	ProjectStatusPlanning   ProjectStatus = "PLANNING"
// 	ProjectStatusActive     ProjectStatus = "ACTIVE"
// 	ProjectStatusOnHold     ProjectStatus = "ON_HOLD"
// 	ProjectStatusCompleted  ProjectStatus = "COMPLETED"
// 	ProjectStatusCancelled  ProjectStatus = "CANCELLED"
// )

// type ActivityType string

// const (
// 	ActivityTypeTaskUpdate  ActivityType = "TASK_UPDATE"
// 	ActivityTypeUserComment ActivityType = "USER_COMMENT"
// )

// type Workspace struct {
// 	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
// 	Name        string    `gorm:"not null" json:"name"`
// 	UrlSlug     string    `gorm:"unique;not null" json:"urlSlug"`
// 	Description *string   `json:"description"`
// 	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
// 	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
// 	// Relationships
// 	// Teams    []Team          `gorm:"foreignKey:WorkspaceId" json:"teams,omitempty"`
// 	Members  []UserWorkspace `gorm:"foreignKey:WorkspaceId" json:"members,omitempty"`
// }

// type Project struct {
// 	ID          uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
// 	Name        string        `gorm:"not null" json:"name"`
// 	Description *string       `json:"description,omitempty"`
// 	StartDate   *time.Time    `json:"startDate,omitempty"`
// 	EndDate     *time.Time    `json:"endDate,omitempty"`
// 	Status      ProjectStatus `gorm:"not null" json:"status"`
// 	CreatedAt   time.Time     `gorm:"not null" json:"createdAt"`
// 	UpdatedAt   time.Time     `gorm:"not null" json:"updatedAt"`
// 	// Relationships
// 	// Teams []ProjectTeam `gorm:"foreignKey:ProjectId" json:"teams,omitempty"`
// 	Tasks []Task        `gorm:"foreignKey:ProjectId" json:"tasks,omitempty"`
// }

// type Team struct {
// 	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
// 	Name        string    `gorm:"not null" json:"name"`
// 	Acronym     string    `gorm:"not null" json:"acronym"`
// 	Description *string   `json:"description,omitempty"`
// 	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
// 	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
// 	WorkspaceId uuid.UUID `gorm:"not null" json:"workspaceId"`
// 	// Relationships
// 	Workspace Workspace     `gorm:"foreignKey:WorkspaceId" json:"workspace"`
// 	Tasks     []Task        `gorm:"foreignKey:TeamId" json:"tasks,omitempty"`
// 	Labels    []TaskLabel   `gorm:"foreignKey:TeamId" json:"labels,omitempty"`
// 	Members   []UserTeam    `gorm:"foreignKey:TeamId" json:"members,omitempty"`
// 	Projects  []ProjectTeam `gorm:"foreignKey:TeamId" json:"projects,omitempty"`
// }

// type UserTeam struct {
// 	UserId    uuid.UUID `gorm:"primaryKey" json:"userId"`
// 	TeamId    uuid.UUID `gorm:"primaryKey" json:"teamId"`
// 	Role      UserRole  `gorm:"not null" json:"role"`
// 	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
// 	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
// 	// Relationships
// 	User User `gorm:"foreignKey:UserId" json:"user"`
// 	Team Team `gorm:"foreignKey:TeamId" json:"team"`
// }

// type ProjectTeam struct {
// 	ProjectId uuid.UUID `gorm:"primaryKey" json:"projectId"`
// 	TeamId    uuid.UUID `gorm:"primaryKey" json:"teamId"`
// 	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
// 	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
// 	// Relationships
// 	Project Project `gorm:"foreignKey:ProjectId" json:"project"`
// 	Team    Team    `gorm:"foreignKey:TeamId" json:"team"`
// }

// type Activity struct {
// 	ID        uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
// 	Content   string       `gorm:"not null" json:"content"`
// 	Type      ActivityType `gorm:"not null" json:"type"`
// 	TaskId    uuid.UUID    `gorm:"not null" json:"taskId"`
// 	CreatedBy *uuid.UUID   `json:"createdBy,omitempty"`
// 	CreatedAt time.Time    `gorm:"not null" json:"createdAt"`
// 	UpdatedAt time.Time    `gorm:"not null" json:"updatedAt"`
// 	// Relationships
// 	Task Task `gorm:"foreignKey:TaskId" json:"task"`
// 	User User `gorm:"foreignKey:CreatedBy" json:"user,omitempty"`
// }
