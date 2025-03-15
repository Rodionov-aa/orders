package entity

import (
	"time"

	"github.com/google/uuid"
)

type Clients struct {
	ID        uuid.UUID
	Inn       string
	IsTrash   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
