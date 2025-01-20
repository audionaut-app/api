package authn

import (
	"context"

	e "encore.app/internal/errs"
	token "encore.app/services/authn/models/token"
	"encore.dev/beta/auth"
)

// CreateCurrentUserToken creates a token for the current user.
//
//encore:api auth method=POST path=/v1/user/me/tokens
func (s *Service) CreateCurrentUserToken(ctx context.Context, params *token.NewToken) (*TokenResponse, error) {
	uid, ok := auth.UserID()
	if !ok {
		return nil, e.UnauthenticatedResponse
	}

	return s.CreateTokenByUserId(ctx, string(uid), params)
}
