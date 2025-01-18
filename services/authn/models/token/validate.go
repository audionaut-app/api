package token

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"errors"
	"time"

	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

type ValidateTokenParams struct {
	Token string `json:"token"`
}

// Validate validates the ValidateTokenParams fields.
func (params *ValidateTokenParams) Validate() error {
	if len(params.Token) == 0 {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "parameter 'token' is required",
		}
	}

	return nil
}

// ValidateToken validates a token and returns the token's user id if successful.
func (m *Model) ValidateToken(ctx context.Context, params *ValidateTokenParams) (string, error) {
	hash := sha256.Sum256([]byte(params.Token))

	row := m.DB.QueryRow(ctx, `
		UPDATE token_
		SET accessed_at = NOW()
		WHERE hash = $1
			AND active = true
			AND (expires_at IS NULL OR expires_at > NOW())
		RETURNING id, user_id, name, active, accessed_at, expires_at, created_at, updated_at, version
	`, hash[:])

	var token Token
	err := row.Scan(&token.Id, &token.UserId, &token.Name, &token.Active, &token.AccessedAt, &token.ExpiresAt, &token.CreatedAt, &token.UpdatedAt, &token.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			dummy := make([]byte, sha256.Size)
			subtle.ConstantTimeCompare(hash[:], dummy)
			return "", &errs.Error{
				Code:    errs.Unauthenticated,
				Message: "invalid or missing authentication credentials",
			}
		default:
			rlog.Error("failed to scan row", "error", err)
			return "", &errs.Error{
				Code:    errs.Internal,
				Message: "the server encountered a problem and could not process your request",
			}
		}
	}

	// If the expires_at field is set and expired, return an error.
	if token.ExpiresAt != nil && token.ExpiresAt.Before(time.Now()) {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "expired authentication credentials",
		}
	}

	return token.UserId, nil
}
