package models

// Vote is a representation of a vote in the database
type Vote struct {
	UserID    uint `json:"user,omitempty" gorm:"unique_index;not null"`
	ArticleID uint `json:"article,omitempty" gorm:"unique_index;not null"`
	Base
}
