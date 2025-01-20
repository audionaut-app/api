package authn

import (
	"context"

	"encore.app/services/authn/models/token"
	"encore.app/services/authz"
	"encore.dev/beta/auth"
)

type AuthParams struct {
	APIKey string `header:"X-API-Key" encore:"sensitive"`
}

//encore:authhandler
func (s *Service) AuthHandler(ctx context.Context, params *AuthParams) (auth.UID, *authz.PermissionsResponse, error) {
	tp := &token.ValidateTokenParams{Token: params.APIKey}
	u, err := s.ValidateToken(ctx, tp)
	if err != nil {
		return "", nil, err
	}

	permissions, err := authz.GetPermissionsByUserId(ctx, u.User.Id)
	if err != nil {
		return "", nil, err
	}

	return auth.UID(u.User.Id), permissions, nil
}
