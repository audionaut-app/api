package authn

import (
	"context"

	e "encore.app/internal/errs"
	"encore.dev/beta/auth"
)

// GetCurrentUser returns the current user.
//
//encore:api auth method=GET path=/v1/user/me
func (s *Service) GetCurrentUser(ctx context.Context) (*UserResponse, error) {
	uid, ok := auth.UserID()
	if !ok {
		return nil, e.UnauthenticatedResponse
	}

	return s.GetUserById(ctx, string(uid))
}
