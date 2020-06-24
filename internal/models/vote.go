package models

import (
	"fmt"

	"github.com/pronuu/lincoln/internal/utils"
)

// Vote is a representation of a vote in the database
type Vote struct {
	UserID    uint `json:"user,omitempty" gorm:"unique_index;not null"`
	ArticleID uint `json:"article,omitempty" gorm:"unique_index;not null"`
	Base
}

// String ...
func (v *Vote) String() string {
	if v == nil {
		return ""
	}
	result := fmt.Sprintf(`
		Vote {
			%s
			%s
			%s
			%s
			%s
		}
	`,
		v.stringID(),
		v.stringUserID(),
		v.stringArticleID(),
		v.stringCreatedAt(),
		v.stringUpdatedAt(),
	)
	return utils.TrimWhitespace(result)
}

func (v *Vote) stringID() string {
	asString := fmt.Sprintf("ID: %d,", v.ID)
	if v.ID == 0 {
		asString = ""
	}
	return asString
}

func (v *Vote) stringUserID() string {
	asString := fmt.Sprintf("UserID: %d,", v.UserID)
	if v.UserID == 0 {
		asString = ""
	}
	return asString
}

func (v *Vote) stringArticleID() string {
	asString := fmt.Sprintf("ArticleID: %d,", v.ArticleID)
	if v.ArticleID == 0 {
		asString = ""
	}
	return asString
}

func (v *Vote) stringCreatedAt() string {
	asString := fmt.Sprintf("CreatedAt: %q,", v.CreatedAt)
	if v.CreatedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (v *Vote) stringUpdatedAt() string {
	asString := fmt.Sprintf("UpdatedAt: %q,", v.UpdatedAt)
	if v.UpdatedAt.IsZero() {
		asString = ""
	}
	return asString
}
