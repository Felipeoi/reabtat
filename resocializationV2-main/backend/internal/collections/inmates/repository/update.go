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

// PUT: atualiza TODOS os campos do registro.
// Retorna o id atualizado. Se não encontrar, retorna pgx.ErrNoRows.
func (r *UpdateRepo) Execute(ctx context.Context, in entity.InmatesFlat, userID int64) (error, context.Context, int) {
	if in.Id == nil {
		return ErrMissingID, ctx, 0
	}

	attorney := pgtype.Text{String: in.Responsible.Attorney, Valid: in.Responsible.Attorney != ""}
	phone := pgtype.Text{String: in.Responsible.Phone, Valid: in.Responsible.Phone != ""}

	var id int
	err := r.pool.QueryRow(ctx, updateInmateSQL,
		*in.Id,           // $1 id
		userID,           // $2 user_id
		in.Custody,       // $3 custody
		in.OriginID,      // $4 origin_id
		in.DestinationID, // $5 destination_id
		attorney,         // $6 responsible_attorney
		phone,            // $7 responsible_phone
	).Scan(&id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// se preferir, mapeie para um erro seu, ex.: ErrInmateNotFound
			return pgx.ErrNoRows, ctx, 0
		}
		return err, ctx, 0
	}
	return nil, ctx, id
}

const updateInmateSQL = `
UPDATE resocialization.inmates
SET
  custody              = $3,
  origin_id            = $4,
  destination_id       = $5,
  responsible_attorney = $6,
  responsible_phone    = $7
WHERE id = $1 AND user_id = $2
RETURNING id;
`
