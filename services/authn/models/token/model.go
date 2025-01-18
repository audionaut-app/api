package token

import (
	"fmt"
	"time"

	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

type Model struct {
	DB *sqldb.Database
}

type Token struct {
	Id         string     `json:"id"`
	Value      TokenValue `json:"value,omitempty" encore:"sensitive"`
	UserId     string     `json:"user_id"`
	Name       string     `json:"name"`
	Active     bool       `json:"active"`
	AccessedAt time.Time  `json:"accessed_at"`
	ExpiresAt  *time.Time `json:"expires_at"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	Version    int32      `json:"-"`
}

type TokenValue struct {
	Plaintext string `json:"plaintext,omitempty" encore:"sensitive"`
	Hash      []byte `json:"-" encore:"sensitive"`
}

type NewToken struct {
	Name      string     `json:"name"`
	ExpiresAt *time.Time `json:"expires_at"`
}

// Validate validates the NewToken fields.
func (params *NewToken) Validate() error {
	if len(params.Name) == 0 {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("parameter 'name' is required"),
		}
	}

	if params.ExpiresAt != nil && params.ExpiresAt.IsZero() {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("parameter 'expires_at' is required"),
		}
	}

	if params.ExpiresAt != nil && params.ExpiresAt.Before(time.Now()) {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("parameter 'expires_at' must be in the future"),
		}
	}

	return nil
}
