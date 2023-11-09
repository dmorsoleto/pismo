package database

import (
	"dmorsoleto/internal/entity"

	"github.com/jmoiron/sqlx"
)

type DatabaseHelper interface {
	InitConnection(databseSettings entity.DatabaseSettings) (*sqlx.DB, error)
	GetConnection() *sqlx.DB
}
