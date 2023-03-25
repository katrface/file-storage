package file_info

import (
	"time"
)

type FileInfo struct {
	Id  uint   `json:"id" gorm:"primaryKey"`
	Url string `json:"url"`
	// Payload   FilePayload `json:"payload,omitempty" gorm:"type:jsonb;default:'{}';not null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// type FilePayload = pgtype.JSONB
