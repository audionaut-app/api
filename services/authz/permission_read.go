package authz

import (
	"context"

	"encore.app/services/authz/models/permission"
)

type PermissionsResponse struct {
	Permissions []permission.Permission `json:"permissions"`
}

// Has checks if the response contains a specific permission.
func (res *PermissionsResponse) Has(permission string) bool {
	for _, p := range res.Permissions {
		if p.String() == permission {
			return true
		}
	}
	return false
}

// GetPermissionsByUserId returns a user's permissions by id.
//
//encore:api private method=GET path=/v1/users/:id/permissions
func (s *Service) GetPermissionsByUserId(ctx context.Context, id string) (*PermissionsResponse, error) {
	permissions, err := s.Permission.GetPermissionsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	return &PermissionsResponse{Permissions: permissions}, nil
}
