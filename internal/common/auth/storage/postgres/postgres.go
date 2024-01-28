package postgres

import (
	"github.com/erupshis/golang-integration-developer-test/internal/common/auth/storage"
	"github.com/erupshis/golang-integration-developer-test/internal/common/db"
	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
)

var (
	_ storage.BaseAuthStorage = (*Postgres)(nil)
)

type Postgres struct {
	*db.Connection

	logger logger.BaseLogger
}

// NewPostgres creates postgresql implementation. Supports migrations and check connection to database.
func NewPostgres(connection *db.Connection, logger logger.BaseLogger) storage.BaseAuthStorage {
	return &Postgres{
		Connection: connection,
		logger:     logger,
	}
}
