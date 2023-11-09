package database

import (
	"dmorsoleto/internal/entity"
	"fmt"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	maxIdleConnections = 2
	connectionLifeSpan = time.Second * 10
	maxConnections     = 10
)

var (
	pool *sqlx.DB
	err  error
)

type databaseHelper struct{}

func NewDatabseHelper() DatabaseHelper {
	return &databaseHelper{}
}

func (databaseHelper) InitConnection(databseSettings entity.DatabaseSettings) (*sqlx.DB, error) {
	settingsDatabase := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s", databseSettings.User, strings.Trim(databseSettings.Pwd, "\n"), databseSettings.Host, databseSettings.Port, databseSettings.DbName, databseSettings.Schema)

	pool, err = sqlx.Connect("pgx", settingsDatabase)
	if err != nil {
		return nil, err
	}

	pool.SetMaxIdleConns(maxIdleConnections)
	pool.SetConnMaxLifetime(connectionLifeSpan)
	pool.SetMaxOpenConns(maxConnections)

	return pool, nil
}

func (databaseHelper) GetConnection() *sqlx.DB {
	return pool
}
