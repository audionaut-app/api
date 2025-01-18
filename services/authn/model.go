package authn

import (
	"encore.app/services/authn/models/token"
	"encore.app/services/authn/models/user"
)

type TokenResponse struct {
	Token token.Token `json:"token"`
}

type UserResponse struct {
	User user.User `json:"user"`
}
