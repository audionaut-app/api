package authn

import (
	"context"

	e "encore.app/internal/errs"
	"encore.app/internal/permission"
	"encore.app/services/authn/models/token"
	"encore.app/services/authz"
	"encore.dev/beta/auth"
)

// CreateTokenByUserId creates a token for the user with "id".
//
//encore:api auth method=POST path=/v1/users/:id/tokens
func (s *Service) CreateTokenByUserId(ctx context.Context, id string, params *token.NewToken) (*TokenResponse, error) {
	uid, ok := auth.UserID()
	if !ok {
		return nil, e.UnauthenticatedResponse
	}

	data, ok := auth.Data().(*authz.PermissionsResponse)
	if !ok {
		return nil, e.UnauthenticatedResponse
	}

	response := func(ctx context.Context, id string) (*TokenResponse, error) {
		token, err := s.Token.CreateTokenByUserId(ctx, id, params)
		if err != nil {
			return nil, err
		}

		return &TokenResponse{Token: *token}, nil
	}

	if data.Has(permission.CreateTokenGlobal) {
		return response(ctx, id)
	}

	if data.Has(permission.CreateTokenSelf) {
		if id == string(uid) {
			return response(ctx, id)
		}
		return nil, e.AccessToResourceDeniedResponse
	}

	return nil, e.UnauthorizedResponse(permission.CreateTokenGlobal, permission.CreateTokenSelf)
}
