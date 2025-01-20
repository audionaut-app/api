package authz

import (
	"context"
	_ "embed"
	"log"

	"encore.app/services/authz/models/permission"
	"encore.dev"
	"encore.dev/rlog"
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("authz", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

//go:embed fixtures.sql
var fixtures string

// Service represents the Encore service application for authorization.
//
//encore:service
type Service struct {
	Permission *permission.Model
}

// NewService creates a new Encore authorization service.
func NewService(db *sqldb.Database) (*Service, error) {
	switch encore.Meta().Environment.Cloud {
	case encore.CloudLocal:
		if _, err := db.Exec(context.Background(), fixtures); err != nil {
			log.Fatalln("failed to add fixtures:", err)
		}
	}

	return &Service{
		Permission: &permission.Model{DB: db},
	}, nil
}

// initService is called by Encore to initialize the service.
func initService() (*Service, error) {
	return NewService(db)
}

// Shutdown is called by Encore to signal the service that it will be shutdown.
func (s *Service) Shutdown(force context.Context) {
	defer rlog.Info("shutdown", "status", "shutdown complete")
}
