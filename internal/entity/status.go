package entity

import "time"

type ObjectStatus struct {
	CreatedAt     time.Time `json:"created_at"`
	LastChangedAt time.Time `json:"last_changed"`

	IsDeleted   bool      `json:"is_deleted"`
	DeletedBy   *User     `json:"deleted_by"`
	DeletedWhen time.Time `json:"deleted_when"`
}
