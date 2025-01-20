package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	e "encore.app/internal/errs"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// GetUserById returns a user by user id.
func (m *Model) GetUserById(ctx context.Context, id string) (*User, error) {
	row := m.DB.QueryRow(ctx, `
		SELECT id, created_at, updated_at, version
		FROM user_
		WHERE id = $1
	`, id)

	var u User
	err := row.Scan(&u.Id, &u.CreatedAt, &u.UpdatedAt, &u.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, &errs.Error{
				Code:    errs.NotFound,
				Message: fmt.Sprintf("user with id '%s' could not be found", id),
			}
		default:
			rlog.Error("the server encountered a problem during token validation", "error", err)
			return nil, e.InternalErrorResponse
		}
	}

	return &u, nil
}
