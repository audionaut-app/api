package authn

import (
	"context"

	"encore.app/services/authn/models/token"
)

// ValidateToken validates a user's token and returns the user's data if successful.
//
//encore:api private method=POST path=/v1/tokens/validate
func (s *Service) ValidateToken(ctx context.Context, params *token.ValidateTokenParams) (*UserResponse, error) {
	id, err := s.Token.ValidateToken(ctx, params)
	if err != nil {
		return nil, err
	}

	u, err := s.User.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserResponse{User: *u}, nil
}
