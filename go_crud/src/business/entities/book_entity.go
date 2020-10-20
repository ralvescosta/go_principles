package entities

import (
	"database/sql"
	"time"
)

// BookEntity ...
type BookEntity struct {
	ID                uint64
	Title             string
	Author            string
	PublishingCompany string
	Edition           uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime
}
