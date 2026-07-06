package repository

import (
	"context"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateRepo struct{ pool *pgxpool.Pool }

func NewCreateRepo(pool *pgxpool.Pool) *CreateRepo { return &CreateRepo{pool: pool} }

func (r *CreateRepo) Execute(ctx context.Context, in entity.InmatesFlat) (error, context.Context, int) {
	var id int

	attorney := pgtype.Text{String: in.Responsible.Attorney, Valid: in.Responsible.Attorney != ""}
	phone := pgtype.Text{String: in.Responsible.Phone, Valid: in.Responsible.Phone != ""}

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err, ctx, 0
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, createInmateSQL,
		in.Custody,
		in.OriginUnitID,
		attorney,
		phone,
		in.UserID,
	).Scan(&id)
	if err != nil {
		return err, ctx, 0
	}

	if err := ReplaceDestinationUnits(ctx, tx, id, in.DestinationUnitIDs); err != nil {
		return err, ctx, 0
	}

	if err := tx.Commit(ctx); err != nil {
		return err, ctx, 0
	}

	return nil, ctx, id
}

var createInmateSQL = `
INSERT INTO resocialization.inmates (
  custody,
  origin_unit_id,
  responsible_attorney,
  responsible_phone,
  user_id
) VALUES ($1,$2,$3,$4,$5)
RETURNING id;
`
