package models

import (
	"fmt"
	"time"

	"github.com/pronuu/lincoln/internal/utils"
)

// Article is a representation of an article in the database
type Article struct {
	Title       string    `json:"title,omitempty" gorm:"not null"`
	URL         string    `json:"url,omitempty" gorm:"not null"`
	Thumbnail   string    `json:"thumbnail,omitempty" gorm:"not null"`
	PublishedAt time.Time `json:"publishedAt,omitempty" gorm:"not null"`
	Votes       []Vote    `json:"votes,omitempty" gorm:"foreignkey:ArticleID"`
	Base
}

// String ...
func (a *Article) String() string {
	if a == nil {
		return ""
	}
	result := fmt.Sprintf(`
		Article{
			%s
			%s
			%s
			%s
			%s
			%s
			%s
		}
	`,
		a.stringID(),
		a.stringTitle(),
		a.stringURL(),
		a.stringThumbnail(),
		a.stringPublishedAt(),
		a.stringCreatedAt(),
		a.stringUpdatedAt(),
	)
	return utils.TrimWhitespace(result)
}

func (a *Article) stringID() string {
	asString := fmt.Sprintf("ID: %d,", a.ID)
	if a.ID == 0 {
		asString = ""
	}
	return asString
}

func (a *Article) stringTitle() string {
	asString := fmt.Sprintf("Title: %q,", a.Title)
	if a.Title == "" {
		asString = ""
	}
	return asString
}

func (a *Article) stringURL() string {
	asString := fmt.Sprintf("URL: %q,", a.URL)
	if a.URL == "" {
		asString = ""
	}
	return asString
}

func (a *Article) stringThumbnail() string {
	asString := fmt.Sprintf("Thumbnail: %q,", a.Thumbnail)
	if a.Thumbnail == "" {
		asString = ""
	}
	return asString
}

func (a *Article) stringPublishedAt() string {
	asString := fmt.Sprintf("PublishedAt: %q,", a.PublishedAt)
	if a.PublishedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (a *Article) stringCreatedAt() string {
	asString := fmt.Sprintf("CreatedAt: %q,", a.CreatedAt)
	if a.CreatedAt.IsZero() {
		asString = ""
	}
	return asString
}

func (a *Article) stringUpdatedAt() string {
	asString := fmt.Sprintf("UpdatedAt: %q,", a.UpdatedAt)
	if a.UpdatedAt.IsZero() {
		asString = ""
	}
	return asString
}
