package models

import "time"

// Base is the attributes needed to turn a model into a table in the database
type Base struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
