package repository

import (
	"context"
	"errors"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrInmateNotFound = errors.New("inmate not found")

type ByIdRepo struct{ pool *pgxpool.Pool }

func NewByIdRepo(pool *pgxpool.Pool) *ByIdRepo { return &ByIdRepo{pool: pool} }

func (r *ByIdRepo) Execute(ctx context.Context, id int, userID int64) (error, context.Context, *entity.InmatesResp) {
	var (
		idDB          int32
		custody       string
		originID      int32
		destinationID int32
		attorney      pgtype.Text
		phone         pgtype.Text
		ownerUserID   int64
		// City fields for origin
		originCityID   int32
		originIBGECode string
		originName     string
		originUFCode   string
		// City fields for destination
		destCityID   int32
		destIBGECode string
		destName     string
		destUFCode   string
	)

	err := r.pool.QueryRow(ctx, getInmateByIDSQL, id, userID).Scan(
		&idDB,
		&custody,
		&originID,
		&destinationID,
		&attorney,
		&phone,
		&ownerUserID,
		// Origin city
		&originCityID,
		&originIBGECode,
		&originName,
		&originUFCode,
		// Destination city
		&destCityID,
		&destIBGECode,
		&destName,
		&destUFCode,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrInmateNotFound, ctx, nil
		}
		return err, ctx, nil
	}

	resp := &entity.InmatesResp{
		Id:            int(idDB),
		Custody:       custody,
		OriginID:      int(originID),
		DestinationID: int(destinationID),
		Origin: &entity.City{
			ID:       int(originCityID),
			IBGECode: originIBGECode,
			Name:     originName,
			UFCode:   originUFCode,
		},
		Destination: &entity.City{
			ID:       int(destCityID),
			IBGECode: destIBGECode,
			Name:     destName,
			UFCode:   destUFCode,
		},
		Responsible: entity.InmatesResponsible{
			Attorney: attorney.String, // vazio se NULL
			Phone:    phone.String,    // vazio se NULL
		},
		UserID: ownerUserID,
	}

	return nil, ctx, resp
}

const getInmateByIDSQL = `
SELECT
  i.id,
  i.custody,
  i.origin_id,
  i.destination_id,
  i.responsible_attorney,
  i.responsible_phone,
  i.user_id,
  oc.id, oc.ibge_code, oc.name, oc.uf_code,
  dc.id, dc.ibge_code, dc.name, dc.uf_code
FROM resocialization.inmates AS i
LEFT JOIN public.cities AS oc ON i.origin_id = oc.id
LEFT JOIN public.cities AS dc ON i.destination_id = dc.id
WHERE i.id = $1 AND i.user_id = $2
`
