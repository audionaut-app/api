package authn

import (
	"context"

	"encore.app/services/authn/models/token"
)

// CreateTokenByUserId creates a token for the user with "id".
//
//encore:api public method=POST path=/v1/users/:id/tokens
func (s *Service) CreateTokenByUserId(ctx context.Context, id string, params *token.NewToken) (*TokenResponse, error) {
	token, err := s.Token.CreateTokenByUserId(ctx, id, params)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{Token: *token}, nil
}
