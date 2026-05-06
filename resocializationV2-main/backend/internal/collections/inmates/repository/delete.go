package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DeleteRepo struct{ pool *pgxpool.Pool }

func NewDeleteRepo(pool *pgxpool.Pool) *DeleteRepo { return &DeleteRepo{pool: pool} }

func (r *DeleteRepo) Execute(ctx context.Context, id int, userID int64) (error, context.Context) {
	var deletedID int64
	err := r.pool.QueryRow(ctx, deleteInmateSQL, id, userID).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Reutilize a mesma variável ErrInmateNotFound definida no seu pacote (GetByID).
			return ErrInmateNotFound, ctx
		}
		return err, ctx
	}
	return nil, ctx
}

var deleteInmateSQL = `
	DELETE FROM resocialization.inmates
	WHERE id = $1 AND user_id = $2
	RETURNING id;
`
