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

	err := r.pool.QueryRow(ctx, createInmateSQL,
		in.Custody,       // custody
		in.OriginID,      // origin_id
		in.DestinationID, // destination_id
		attorney,         // responsible_attorney (NULL se vazio)
		phone,            // responsible_phone (NULL se vazio)
		in.UserID,        // user_id
	).Scan(&id)

	return err, ctx, id
}

var createInmateSQL = `
INSERT INTO resocialization.inmates (
  custody,
  origin_id,
  destination_id,
  responsible_attorney,
  responsible_phone,
  user_id
) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id;
`
