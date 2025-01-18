package user

import (
	"fmt"
	"time"

	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

type Model struct {
	DB *sqldb.Database
}

type User struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Version   int32     `json:"-"`
}

type NewUser struct {
	Id string `json:"id"`
}

// Validate validates the NewUser fields.
func (params *NewUser) Validate() error {
	if len(params.Id) == 0 {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("parameter 'id' is required"),
		}
	}

	return nil
}
