package permission

import (
	"fmt"
	"time"

	"encore.dev/storage/sqldb"
)

type Model struct {
	DB *sqldb.Database
}

type Permission struct {
	OperationId string    `json:"operation_id"`
	ResourceId  string    `json:"resource_id"`
	ContextId   string    `json:"context_id"`
	Description *string   `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Version     int32     `json:"-"`
}

// String returns a string representation of Permission.
func (p *Permission) String() string {
	return fmt.Sprintf("%s:%s:%s", p.OperationId, p.ResourceId, p.ContextId)
}
