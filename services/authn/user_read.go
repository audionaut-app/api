package authn

import (
	"context"
)

// GetUserById returns a user by user id.
//
//encore:api private method=GET path=/v1/users/:id
func (s *Service) GetUserById(ctx context.Context, id string) (*UserResponse, error) {
	u, err := s.User.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserResponse{User: *u}, nil
}
