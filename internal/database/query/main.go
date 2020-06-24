package query

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Query ...
type Query struct {
	Order      string
	Offset     int
	Limit      int
	Parameters []Parameter
}

// Parameter ...
type Parameter interface{}

// Execute will take in a db instance and run the query on it
func (q *Query) Execute(db *gorm.DB) (*sql.Rows, error) {
	db = db.Order(q.Order).Limit(q.Limit).Offset(q.Offset)
	for _, parameter := range q.Parameters {
		db = db.Where(parameter)
	}
	return db.Rows()
}
