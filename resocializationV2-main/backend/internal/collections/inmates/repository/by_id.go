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
		idDB             int32
		custody          string
		originUnitID     int32
		attorney         pgtype.Text
		phone            pgtype.Text
		ownerUserID      int64
		originUnitIDJoin int32
		originUnitName   string
		originUnitUF     string
	)

	err := r.pool.QueryRow(ctx, getInmateByIDSQL, id, userID).Scan(
		&idDB,
		&custody,
		&originUnitID,
		&attorney,
		&phone,
		&ownerUserID,
		&originUnitIDJoin,
		&originUnitName,
		&originUnitUF,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrInmateNotFound, ctx, nil
		}
		return err, ctx, nil
	}

	destMap, err := LoadDestinationUnitsMap(ctx, r.pool, []int{int(idDB)})
	if err != nil {
		return err, ctx, nil
	}

	destUnits := destMap[int(idDB)]
	if destUnits == nil {
		destUnits = []entity.PrisonUnit{}
	}

	resp := &entity.InmatesResp{
		Id:                 int(idDB),
		Custody:            custody,
		OriginUnitID:       int(originUnitID),
		DestinationUnitIDs: DestinationUnitIDs(destUnits),
		DestinationUnits:   destUnits,
		OriginUnit: &entity.PrisonUnit{
			ID:     int(originUnitIDJoin),
			Name:   originUnitName,
			UFCode: originUnitUF,
		},
		Responsible: entity.InmatesResponsible{
			Attorney: attorney.String,
			Phone:    phone.String,
		},
		UserID: ownerUserID,
	}

	return nil, ctx, resp
}

const getInmateByIDSQL = `
SELECT
  i.id,
  i.custody,
  i.origin_unit_id,
  i.responsible_attorney,
  i.responsible_phone,
  i.user_id,
  ou.id, ou.name, ou.uf_code
FROM resocialization.inmates AS i
LEFT JOIN public.prison_units AS ou ON i.origin_unit_id = ou.id
WHERE i.id = $1 AND i.user_id = $2
`
