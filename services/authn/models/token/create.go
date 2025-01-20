package token

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"

	e "encore.app/internal/errs"
	"encore.app/internal/xid"
	"encore.dev/rlog"
)

// createToken creates a cryptographically secure token.
func createToken() (string, []byte, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", nil, err
	}

	// Encode bytes to base32 string.
	token := xid.Token{}.Prefix() + "_" + base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)

	// Generate SHA256 hash of the token.
	hash := sha256.Sum256([]byte(token))
	return token, hash[:], nil
}

// CreateTokenByUserId creates a token for the user with "id".
func (m *Model) CreateTokenByUserId(ctx context.Context, id string, params *NewToken) (*Token, error) {
	plaintext, hash, err := createToken()
	if err != nil {
		rlog.Error("failed to create token", "error", err)
		return nil, e.InternalErrorResponse
	}

	row := m.DB.QueryRow(ctx, `
		INSERT INTO token_ (id, hash, user_id, name, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, name, active, accessed_at, expires_at, created_at, updated_at, version
	`, xid.New[xid.Token]().String(), hash, id, params.Name, params.ExpiresAt)

	token := Token{
		Value: TokenValue{
			Plaintext: plaintext,
			Hash:      hash,
		},
	}

	err = row.Scan(&token.Id, &token.UserId, &token.Name, &token.Active, &token.AccessedAt, &token.ExpiresAt, &token.CreatedAt, &token.UpdatedAt, &token.Version)
	if err != nil {
		rlog.Error("failed to scan row", "error", err)
		return nil, e.InternalErrorResponse
	}

	return &token, nil
}
