package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pronuu/lincoln/internal/utils"
)

// VoteArray ...
type VoteArray []*Vote

// String ...
func (voteArray *VoteArray) String() string {
	voteIDs := make([]string, 0)
	for _, vote := range *voteArray {
		voteIDs = append(voteIDs, fmt.Sprintf("Vote{ID: %s}", strconv.Itoa(int(vote.ID))))
	}
	return fmt.Sprintf("[]Vote{%s}", strings.Join(voteIDs, ","))
}

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
