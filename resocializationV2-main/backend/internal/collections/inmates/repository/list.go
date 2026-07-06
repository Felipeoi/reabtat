package repository

import (
	"context"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ListRepo struct{ pool *pgxpool.Pool }

func NewListRepo(pool *pgxpool.Pool) *ListRepo { return &ListRepo{pool: pool} }

func (r *ListRepo) Execute(ctx context.Context, limit, offset int, userID int64) (error, context.Context, []entity.InmatesList) {
	if limit <= 0 || limit > 1000 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	rows, err := r.pool.Query(ctx, listInmatesSQL, userID, limit, offset)
	if err != nil {
		return err, ctx, nil
	}
	defer rows.Close()

	out := make([]entity.InmatesList, 0, limit)
	inmateIDs := make([]int, 0, limit)

	for rows.Next() {
		var (
			id               int32
			custody          string
			originUnitID     int32
			originUnitIDJoin int32
			originUnitName   string
			originUnitUF     string
		)

		if err := rows.Scan(
			&id,
			&custody,
			&originUnitID,
			&originUnitIDJoin,
			&originUnitName,
			&originUnitUF,
		); err != nil {
			return err, ctx, nil
		}

		inmateIDs = append(inmateIDs, int(id))
		out = append(out, entity.InmatesList{
			Id:           int(id),
			Custody:      custody,
			OriginUnitID: int(originUnitID),
			OriginUnit: &entity.PrisonUnit{
				ID:     int(originUnitIDJoin),
				Name:   originUnitName,
				UFCode: originUnitUF,
			},
		})
	}

	if err := rows.Err(); err != nil {
		return err, ctx, nil
	}

	destMap, err := LoadDestinationUnitsMap(ctx, r.pool, inmateIDs)
	if err != nil {
		return err, ctx, nil
	}

	for i := range out {
		destUnits := destMap[out[i].Id]
		if destUnits == nil {
			destUnits = []entity.PrisonUnit{}
		}
		out[i].DestinationUnits = destUnits
		out[i].DestinationUnitIDs = DestinationUnitIDs(destUnits)
	}

	return nil, ctx, out
}

var listInmatesSQL = `
	SELECT
	  i.id,
	  i.custody,
	  i.origin_unit_id,
	  ou.id, ou.name, ou.uf_code
	FROM resocialization.inmates AS i
	LEFT JOIN public.prison_units AS ou ON i.origin_unit_id = ou.id
	WHERE i.user_id = $1
	ORDER BY i.id DESC
	LIMIT $2 OFFSET $3;
`
