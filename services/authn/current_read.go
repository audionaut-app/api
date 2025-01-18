package authn

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

// GetCurrentUser returns the current user.
//
//encore:api auth method=GET path=/v1/user/me
func (s *Service) GetCurrentUser(ctx context.Context) (*UserResponse, error) {
	uid, ok := auth.UserID()
	if !ok {
		return nil, &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid or missing authentication credentials",
		}
	}

	return s.GetUserById(ctx, string(uid))
}
