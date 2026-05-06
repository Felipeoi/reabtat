package pgxerr

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// IsUndefinedRelation indica tabela/schema inexistente (SQLSTATE 42P01).
func IsUndefinedRelation(err error) bool {
	var e *pgconn.PgError
	return errors.As(err, &e) && e.Code == "42P01"
}
