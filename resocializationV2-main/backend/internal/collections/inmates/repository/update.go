package repository

import (
	"context"
	"errors"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrMissingID = errors.New("inmate id is required")

type UpdateRepo struct{ pool *pgxpool.Pool }

func NewUpdateRepo(pool *pgxpool.Pool) *UpdateRepo { return &UpdateRepo{pool: pool} }

func (r *UpdateRepo) Execute(ctx context.Context, in entity.InmatesFlat, userID int64) (error, context.Context, int) {
	if in.Id == nil {
		return ErrMissingID, ctx, 0
	}

	attorney := pgtype.Text{String: in.Responsible.Attorney, Valid: in.Responsible.Attorney != ""}
	phone := pgtype.Text{String: in.Responsible.Phone, Valid: in.Responsible.Phone != ""}

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err, ctx, 0
	}
	defer tx.Rollback(ctx)

	var id int
	err = tx.QueryRow(ctx, updateInmateSQL,
		*in.Id,
		userID,
		in.Custody,
		in.OriginUnitID,
		attorney,
		phone,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgx.ErrNoRows, ctx, 0
		}
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

const updateInmateSQL = `
UPDATE resocialization.inmates
SET
  custody                = $3,
  origin_unit_id         = $4,
  responsible_attorney   = $5,
  responsible_phone      = $6
WHERE id = $1 AND user_id = $2
RETURNING id;
`
