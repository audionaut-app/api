package error

import (
	"fmt"
	"strings"

	"encore.dev/beta/errs"
)

var InternalErrorResponse = &errs.Error{
	Code:    errs.Internal,
	Message: "the server encountered a problem and could not process your request",
}

var UnauthenticatedResponse = &errs.Error{
	Code:    errs.Unauthenticated,
	Message: "invalid or missing authentication credentials",
}

var AccessToResourceDeniedResponse = &errs.Error{
	Code:    errs.PermissionDenied,
	Message: "your user account does not have access to this resource",
}

func UnauthorizedResponse(permissions ...string) error {
	return &errs.Error{
		Code:    errs.PermissionDenied,
		Message: fmt.Sprintf("your user account requires one of the following permissions to access this resource: %s", strings.Join(permissions, ", ")),
	}
}
