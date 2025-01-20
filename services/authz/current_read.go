package authz

import (
	"context"

	e "encore.app/internal/errs"
	"encore.dev/beta/auth"
)

// GetPermissionsByUserId returns the current user's permissions.
//
//encore:api auth method=GET path=/v1/user/me/permissions
func (s *Service) GetCurrentUserPermissions(ctx context.Context) (*PermissionsResponse, error) {
	uid, ok := auth.UserID()
	if !ok {
		return nil, e.UnauthenticatedResponse
	}

	return s.GetPermissionsByUserId(ctx, string(uid))
}
